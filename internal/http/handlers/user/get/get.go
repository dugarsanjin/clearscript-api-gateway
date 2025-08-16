package get

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

// Response todo add status and error fields
type Response struct {
	ID          string       `json:"id"`
	FullName    string       `json:"fullName"`
	Email       string       `json:"email"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Retrieve user information by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/users/{id} [get]
func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.get.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		userId := chi.URLParam(r, "id")
		if userId == "" {
			log.Info("no id provided")

			render.JSON(w, r, map[string]string{"error": "id is required"})
			return
		}

		// Mock response data
		response := getMockResponse(userId)

		log.Info("user retrieved", slog.String("user_id", userId))
		render.JSON(w, r, response)
	}
}

func getMockResponse(userId string) Response {
	return Response{
		ID:       userId,
		FullName: "John Doe",
		Email:    "john.doe@example.com",
		Permissions: []Permission{
			{
				Name:    "admin",
				Actions: []string{"read", "write", "delete"},
			},
			{
				Name:    "user",
				Actions: []string{"read"},
			},
		},
	}
}
