package storage

import (
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MustInit -- FIRMA ACTUALIZADA
func MustInit(dsn string, logger *slog.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Usa el logger inyectado en lugar del `log` global
		logger.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	return db
}
