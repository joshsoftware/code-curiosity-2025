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
	CreateUserGoalInProgress(w http.ResponseWriter, r *http.Request)
	ResetUserCurrentGoalStatus(w http.ResponseWriter, r *http.Request)
	GetUserCurrentGoalStatus(w http.ResponseWriter, r *http.Request)
	FetchUserMonthlyGoalSummary(w http.ResponseWriter, r *http.Request)
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
		slog.Error("error fetching goal levels", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "goal levels fetched successfully", gaols)
}

func (h *handler) CreateUserGoalInProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	var userSelecetdGoal CreateUserGoalRequest
	err := json.NewDecoder(r.Body).Decode(&userSelecetdGoal)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	userGoal, err := h.goalService.CreateUserGoalInProgress(ctx, userSelecetdGoal, userId)
	if err != nil {
		slog.Error("failed to create user goal status", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "Goal created successfully", userGoal)
}

func (h *handler) ResetUserCurrentGoalStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userResetGoalStatus, err := h.goalService.ResetUserCurrentGoalStatus(ctx, userId)
	if err != nil {
		slog.Error("error resetting user current goal status", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user current goal status reset successfully", userResetGoalStatus)
}

func (h *handler) GetUserCurrentGoalStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userCurrentGoalStatus, err := h.goalService.GetUserCurrentGoalStatus(ctx, userId)
	if err != nil {
		slog.Error("error getting current goal status for user", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user current goal status fetched successfully", userCurrentGoalStatus)
}

func (h *handler) FetchUserMonthlyGoalSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdCtxVal := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdCtxVal.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userMonthlyGoalSummary, err := h.goalService.FetchUserGoalSummary(ctx, userId)
	if err != nil {
		slog.Error("error etching user monthly goal summary", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user monthly goal summary fetched successfully", userMonthlyGoalSummary)
}
