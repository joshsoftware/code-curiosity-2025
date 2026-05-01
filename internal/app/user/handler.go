package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	userService Service
}

type Handler interface {
	UpdateUserEmail(w http.ResponseWriter, r *http.Request)
	GetContributionScores(w http.ResponseWriter, r *http.Request)
	UpdateContributionScore(w http.ResponseWriter, r *http.Request)
	GetLeaderboard(w http.ResponseWriter, r *http.Request)
	UpdateUserBlocked(w http.ResponseWriter, r *http.Request)
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

func (h *handler) GetContributionScores(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contributionScores, err := h.userService.GetContributionScores(ctx)
	if err != nil {
		slog.Error("failed to get contribution scores", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "contribution scores fetched successfully", contributionScores)
}

func (h *handler) UpdateContributionScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contributionScoreId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidQueryParams.Error(), nil)
		return
	}

	var requestBody ConfigureContributionScoreRequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	err = h.userService.UpdateContributionScore(ctx, contributionScoreId, requestBody.Score)
	if err != nil {
		slog.Error("failed to update contribution score", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "contribution score updated successfully", nil)
}

func (h *handler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	leaderboard, err := h.userService.GetLeaderboard(ctx)
	if err != nil {
		slog.Error("failed to get leaderboard", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "leaderboard fetched successfully", leaderboard)
}

func (h *handler) UpdateUserBlocked(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidQueryParams.Error(), nil)
		return
	}

	var requestBody BlockUserRequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	err = h.userService.UpdateUserBlocked(ctx, userId, requestBody.IsBlocked)
	if err != nil {
		slog.Error("failed to update user blocked status", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user blocked status updated successfully", nil)
}
