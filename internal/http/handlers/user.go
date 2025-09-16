package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	log *slog.Logger
}

// UserGetResponse todo add status and error fields
type UserGetResponse struct {
	ID          string       `json:"id"`
	FullName    string       `json:"fullName"`
	Email       string       `json:"email"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

func NewUserHandler(log *slog.Logger) *UserHandler {
	return &UserHandler{log: log}
}

// Get godoc
// @Summary      Get user by ID
// @Description  Retrieve user information by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  UserGetResponse
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.Get"

		log := h.log.With(
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

func getMockResponse(userId string) UserGetResponse {
	return UserGetResponse{
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
