package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	userService Service
}

type Handler interface {
	UpdateUserEmail(w http.ResponseWriter, r *http.Request)
	GetAllUsersRank(w http.ResponseWriter, r *http.Request)
	GetCurrentUserRank(w http.ResponseWriter, r *http.Request)
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

func (h *handler) GetAllUsersRank(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	leaderboard, err := h.userService.GetAllUsersRank(ctx)
	if err != nil {
		slog.Error("failed to get all users rank", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "leaderboard fetched successfully", leaderboard)
}

func (h *handler) GetCurrentUserRank(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	currentUserRank, err := h.userService.GetCurrentUserRank(ctx)
	if err != nil {
		slog.Error("failed to get current user rank", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "current user rank fetched successfully", currentUserRank)
}
