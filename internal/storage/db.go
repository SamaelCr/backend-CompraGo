package storage

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInit(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	return db
}