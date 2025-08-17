// Package main provides the entry point for the ClearScript API Gateway.
//
// @title           ClearScript API Gateway
// @version         1.0
// @description     This is a Go-based API gateway for ClearScript application.
// @termsOfService  http://swagger.io/terms/
//
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
//
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
//
// @BasePath  /clearscript-api-gateway
//
// @schemes   http https
package main

import (
	_ "clearscript-api-gateway/docs"
	"clearscript-api-gateway/internal/config"
	contentGet "clearscript-api-gateway/internal/http/handlers/content/get"
	"clearscript-api-gateway/internal/http/handlers/health"
	userGet "clearscript-api-gateway/internal/http/handlers/user/get"
	mwLogger "clearscript-api-gateway/internal/http/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func main() {

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting clearscript-api-gateway", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/clearscript-api-gateway", func(r chi.Router) {
		// Health check
		r.Get("/health", health.New(log))

		// API v1
		r.Get("/api/v1/users/{id}", userGet.New(log))
		r.Get("/api/v1/lessons/{id}", contentGet.New(log))

		// Swagger documentation
		r.Get("/swagger/*", httpSwagger.Handler())
	})

	log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", slog.String("address", cfg.HTTPServer.Address))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
