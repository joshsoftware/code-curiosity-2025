package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	userService Service
}

type Handler interface {
	UpdateUserEmail(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewHandler(userService Service) Handler {
	return &handler{
		userService: userService,
	}
}

func (h *handler) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestBody Email
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	err = h.userService.UpdateUserEmail(ctx, requestBody.Email)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "email updated successfully", nil)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	val := ctx.Value(middleware.UserIdKey)

	userID := val.(int)

	user, err := h.userService.SoftDeleteUser(ctx, userID)
	if err != nil {
		slog.Error("failed to softdelete user", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user scheduled for deletion", user)

}
