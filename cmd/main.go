package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
	"ugly-friend/config"
	"ugly-friend/core"
	"ugly-friend/handler"
	"ugly-friend/middleware"
	"ugly-friend/routes"

	"github.com/go-chi/chi/v5"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	log.Println("⭐️⭐️⭐️ Starting the project...")

	cfg, err := config.MustLoad()
	if err != nil {
		log.Print(fmt.Errorf("failed to load config: %w", err))
	}

	log := setupLogger(cfg.Deploy)

	core := core.InitCore(&cfg.Storage)

	handler := handler.NewHandler(core)

	router := middleware.SetupRouter()
	routes.SetupRoutes(router, handler)

	log.Info("Init server...")

	startRest(cfg, router)
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

// Starts the REST API
func startRest(cfg *config.Config, router *chi.Mux) {
	serverStartURI := fmt.Sprintf("%s:%d",
		cfg.Server.Bind, cfg.Server.Port)

	listener, err := net.Listen("tcp", serverStartURI)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Minute,
		WriteTimeout: 15 * time.Minute,
		IdleTimeout:  15 * time.Minute,
	}

	log.Println("✅ Application initialized and started: ", serverStartURI)

	if err := httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Println("Server shutdown")
		default:
			log.Fatal(err)
		}
	}
}
