package goal

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	goalService Service
}

type Handler interface {
	ListGoalLevels(w http.ResponseWriter, r *http.Request)
	GetUserActiveGoalLevel(w http.ResponseWriter, r *http.Request)
	CreateCustomGoalLevelTarget(w http.ResponseWriter, r *http.Request)
	ListUserGoalLevelProgress(w http.ResponseWriter, r *http.Request)
}

func NewHandler(goalService Service) Handler {
	return &handler{
		goalService: goalService,
	}
}

func (h *handler) ListGoalLevels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	gaols, err := h.goalService.ListGoalLevels(ctx)
	if err != nil {
		slog.Error("error fetching users conributed repos", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "goal levels fetched successfully", gaols)
}

func (h *handler) GetUserActiveGoalLevel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userGoalLevel, err := h.goalService.GetUserActiveGoalLevel(ctx, userId)
	if err != nil {
		slog.Error("error fetching users active goal level", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user active goal level fetched successfully", userGoalLevel)
}

func (h *handler) CreateCustomGoalLevelTarget(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	var customGoalLevelTarget []CustomGoalLevelTarget
	err := json.NewDecoder(r.Body).Decode(&customGoalLevelTarget)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	createdCustomGoalLevelTargets, err := h.goalService.CreateCustomGoalLevelTarget(ctx, userId, customGoalLevelTarget)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "custom goal level targets created successfully", createdCustomGoalLevelTargets)
}

func (h *handler) ListUserGoalLevelProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userGoalLevelProgress, err := h.goalService.ListUserGoalLevelProgress(ctx, userId)
	if err != nil {
		slog.Error("error failed to fetch user goal level progress", "error", err)
		response.WriteJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user goal level progress fetched successfully", userGoalLevelProgress)

}
