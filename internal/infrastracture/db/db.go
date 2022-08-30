package db

import (
	"content-editor/internal/domain/models"
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func DBConnect(cfg *DBConfigSection) (db *gorm.DB, sqlDb *sql.DB, err error) {
	p := postgres.Config{
		DSN:                  cfg.DSN(),
		PreferSimpleProtocol: true,
	}

	db, err = gorm.Open(postgres.New(p), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, nil, err
	}
	sqlDB.SetConnMaxLifetime(time.Second * 20)
	db.AutoMigrate(models.Catalogue{}, models.Item{}, models.Manufacturer{})

	return db, sqlDB, nil
}
