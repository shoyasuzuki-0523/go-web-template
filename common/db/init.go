package db

import (
	"kakeru-pro-web/common/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {
	// gorm の log を zap 経由で出力する（local 環境以外）
	gormLogger := logger.Default

	db, err := gorm.Open(postgres.Open(config.Config.DB_DSN), &gorm.Config{Logger: gormLogger})
	if err != nil {
		panic("Error DB init.")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 発行した SQL をログ出力する（local 環境のみ）
	if config.Env.IsLocal() {
		db.Logger = db.Logger.LogMode(logger.Info)
	}

	return db
}
