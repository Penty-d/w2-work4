package db

import (
	"w2-work4/internal/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open((cfg.DB.DSN())), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
