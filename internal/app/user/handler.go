package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/goal"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	userService Service
}

type Handler interface {
	UpdateUserEmail(w http.ResponseWriter, r *http.Request)
	UpdateCurrentActiveGoalId(w http.ResponseWriter, r *http.Request)
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

func (h *handler) UpdateCurrentActiveGoalId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody goal.GoalLevel

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	slog.Info(reqBody.LevelName)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	goalId, err := h.userService.UpdateCurrentActiveGoalId(ctx, reqBody.LevelName)
	if err != nil {
		slog.Error("[user handler] Failed to update current active goal id", "error", err)
		status, errMsg := apperrors.MapError(err)
		response.WriteJson(w, status, errMsg, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "Goal updated successfully", goal.GoalId{GoalId: goalId})
}
