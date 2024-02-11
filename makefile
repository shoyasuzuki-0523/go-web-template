DB_PORT := 5432
APP_PORT := 8080
ifneq (,$(wildcard ./.env))
	include .env
endif

DB_URL := postgres://postgres:postgres@localhost:${DB_PORT}/postgres?sslmode=disable

.PHONY: migrate-create-ddl
migrate-create-ddl:
	docker run --rm -v `pwd`/app/common/migrations/ddl:/migrations --network=host migrate/migrate -path=/migrations/ -database ${DB_URL} create --ext sql --dir migrations $(NAME)

.PHONY: migrate-create-seed
migrate-create-seed:
	docker run --rm -v `pwd`/app/common/migrations/seed/local:/migrations --network=host migrate/migrate -path=/migrations/ -database ${DB_URL} create --ext sql --dir migrations $(NAME)
