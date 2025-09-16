package handlers

import (
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
	Version   string `json:"version"`
}

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.health.New"

		log.Debug("health check requested", slog.String("op", op))

		// Читаем версию из переменной окружения
		version := os.Getenv("APP_VERSION")
		if version == "" {
			version = "dev"
		}

		response := Response{
			Status:    "ok",
			Timestamp: time.Now().Format(time.RFC3339),
			Service:   "clearscript-api-gateway",
			Version:   version,
		}

		render.JSON(w, r, response)
	}
}
