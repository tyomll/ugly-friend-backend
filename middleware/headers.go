package middleware

import (
	"log"
	"net/http"
	"time"
	"ugly-friend/config"
	"ugly-friend/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter() *chi.Mux {
	router := chi.NewMux()
	cfg, err := config.MustLoad()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
	}

	corsOptions := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Middleware setup
	router.Use(corsOptions)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(func(next http.Handler) http.Handler {
		return utils.JWTMiddleware(next, []byte(cfg.JWT.SecretKey), []string{"/login", "/register"})
	})
	return router
}
