package repository

import (
	"context"
	"fmt"
	"log/slog"
	"ugly-friend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormPostgresDB(ctx context.Context, storage config.Storage, log *slog.Logger) (*gorm.DB, error) {
	log.Debug("Initialize db connection...", "dsn", fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s",
		storage.Host, storage.Username, storage.DBName, storage.Port, storage.SSLMode))

	// Connection pull creation
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		storage.Host, storage.Username, storage.Password, storage.DBName, storage.Port, storage.SSLMode)

	// Open connection pull
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Error opening gorm connection", "error", err)
		return nil, err
	}

	// Ping the connection
	conn, err := db.DB()
	if err != nil {
		log.Error("Error getting DB connection", "error", err)
		return nil, err
	}

	err = conn.PingContext(ctx)
	if err != nil {
		log.Error("DB Connection lost", "error", err)
		return nil, err
	}

	log.Info("Successfully connected to database", "dbname", storage.DBName)

	return db, nil
}
