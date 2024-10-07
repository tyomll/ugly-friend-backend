package core

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"ugly-friend/config"
	"ugly-friend/migrations"
	"ugly-friend/pgxpools"
	"ugly-friend/repository"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type UglyFriendCore struct {
	Methods *repository.Repository
	Logger  *slog.Logger
}

func InitCore(storage *config.Storage) *UglyFriendCore {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Print(fmt.Errorf("failed to load config: %w", err))
	}
	log := setupLogger(cfg.Deploy)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := repository.NewGormPostgresDB(ctx, cfg.Storage, log)
	if err != nil {
		log.Error("💔 DB connection error: %v", slog.Any("error", err))
	}

	err = migrations.MigrateTables(db, log)
	if err != nil {
		log.Error("💔 Migrations error: %v", slog.Any("error", err))
	}

	pool := pgxpools.ConnectDB(&pgxpools.ConfigConnectPgxPool{
		Host:     storage.Host,
		Port:     storage.Port,
		User:     storage.Username,
		Password: storage.Password,
		Name:     storage.DBName,
		SSLMode:  storage.SSLMode,
	})

	methods := repository.NewRepository(pool)

	return &UglyFriendCore{Logger: log, Methods: methods}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
