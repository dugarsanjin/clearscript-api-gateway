package get

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"time"
)

type Response struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Letter      Letter `json:"letter"`
	Description string `json:"description"`
	Author      Author `json:"author"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Letter struct {
	Initial         string `json:"initial"`
	Medial          string `json:"medial"`
	Final           string `json:"final"`
	Transliteration string `json:"transliteration"`
}

type Author struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

// GetLesson godoc
// @Summary      Get lesson by ID
// @Description  Retrieve lesson content by lesson ID
// @Tags         lessons
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Lesson ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/lessons/{id} [get]
func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.content.get.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		contentId := chi.URLParam(r, "id")
		if contentId == "" {
			log.Info("no id provided")

			render.JSON(w, r, map[string]string{"error": "id is required"})
			return
		}

		// Mock response data
		response := getMockResponse(contentId)

		log.Info("content retrieved", slog.String("content_id", contentId))
		render.JSON(w, r, response)
	}
}

func getMockResponse(contentId string) Response {
	now := time.Now().Format(time.RFC3339)

	return Response{
		ID:    contentId,
		Title: "Sample ClearScript Content",
		Letter: Letter{
			Initial:         "𐤀",
			Medial:          "𐤀",
			Final:           "𐤀",
			Transliteration: "aleph",
		},
		Description: "This is a sample content description for ClearScript learning materials",
		Author: Author{
			ID:       "550e8400-e29b-41d4-a716-446655440000",
			FullName: "Dr. Jane Smith",
			Email:    "jane.smith@clearscript.edu",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}
