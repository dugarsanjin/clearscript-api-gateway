package handlers

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type CompareHandler struct {
	log *slog.Logger
}

type CompareResponse struct {
	Similarity float64 `json:"similarity"`
}

func NewCompareHandler(log *slog.Logger) *CompareHandler {
	return &CompareHandler{log: log}
}

// Post godoc
// @Summary      Compare input file with template
// @Description  Compare the uploaded input file against the provided template and return similarity score
// @Tags         compare
// @Accept       multipart/form-data
// @Produce      json
// @Param        template  formData  file  true  "Template file"
// @Param        input     formData  file  true  "Input file to compare"
// @Success      200  {object}  CompareResponse
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/compare [post]
func (h *CompareHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.compare.Post"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Warn("failed to parse multipart form", slog.String("error", err.Error()))

			render.JSON(w, r, map[string]string{"error": "failed to parse form"})
			return
		}

		template, _, err := r.FormFile("template")
		if err != nil {
			log.Warn("template file is required", slog.String("error", err.Error()))

			render.JSON(w, r, map[string]string{"error": "template file is required"})
			return
		}
		defer template.Close()

		inputFile, _, err := r.FormFile("input")
		if err != nil {
			log.Warn("input file is required", slog.String("error", err.Error()))

			render.JSON(w, r, map[string]string{"error": "input file is required"})
			return
		}
		defer inputFile.Close()

		// todo: validate files (type, size, etc.)
		log.Debug("files received")

		response := getCompareMockResponse()
		log.Info("compare completed", slog.Float64("similarity", response.Similarity))

		render.JSON(w, r, response)
	}
}

func getCompareMockResponse() CompareResponse {
	return CompareResponse{
		Similarity: 0.85,
	}
}
