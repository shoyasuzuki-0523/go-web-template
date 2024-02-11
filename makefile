DB_PORT := 5432
APP_PORT := 8080
ifneq (,$(wildcard ./.env))
	include .env
endif

DB_URL := postgres://postgres:postgres@localhost:${DB_PORT}/postgres?sslmode=disable

.PHONY: init
init:
	@make up
	sleep 5
	@make migrate-up
	@make logs

.PHONY: clean
clean:
	docker compose stop
	docker compose down

.PHONY: up
up:
	docker compose up -d

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: migrate-create-ddl
migrate-create-ddl:
	docker run --rm -v `pwd`/common/migrations/ddl:/migrations --network=host migrate/migrate -path=/migrations/ -database ${DB_URL} create --ext sql --dir migrations $(NAME)

.PHONY: migrate-create-seed
migrate-create-seed:
	docker run --rm -v `pwd`/common/migrations/seeds/local:/migrations --network=host migrate/migrate -path=/migrations/ -database ${DB_URL} create --ext sql --dir migrations $(NAME)

.PHONY: migrate-up
migrate-up:
	docker run --rm -v `pwd`/common/migrations/ddl:/migrations --network=host migrate/migrate -path=/migrations/ -database "${DB_URL}&x-migrations-table=ddl_migrations" up
	# docker run --rm -v `pwd`/common/migrations/seeds/local:/migrations --network=host migrate/migrate -path=/migrations/ -database "${DB_URL}&x-migrations-table=seed_migrations" up

.PHONY: migrate-down
migrate-down:
	docker run --rm -v `pwd`/common/migrations/ddl:/migrations --network=host migrate/migrate -path=/migrations/ -database "${DB_URL}&x-migrations-table=ddl_migrations" down -all
	# docker run --rm -v `pwd`/common/migrations/seeds/local:/migrations --network=host migrate/migrate -path=/migrations/ -database "${DB_URL}&x-migrations-table=seed_migrations" down -all
