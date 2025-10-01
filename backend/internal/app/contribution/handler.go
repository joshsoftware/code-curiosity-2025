package contribution

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/utils"
)

type handler struct {
	contributionService Service
}

type Handler interface {
	FetchUserContributions(w http.ResponseWriter, r *http.Request)
	ListMonthlyContributionSummary(w http.ResponseWriter, r *http.Request)
	ListAllContributionTypes(w http.ResponseWriter, r *http.Request)
	ConfigureContributionTypeScore(w http.ResponseWriter, r *http.Request)
}

func NewHandler(contributionService Service) Handler {
	return &handler{
		contributionService: contributionService,
	}
}

func (h *handler) FetchUserContributions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdValue := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userContributions, err := h.contributionService.FetchUserContributions(ctx, userId)
	if err != nil {
		slog.Error("error fetching user contributions", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user contributions fetched successfully", userContributions)
}

func (h *handler) ListMonthlyContributionSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdValue := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	yearVal := r.URL.Query().Get("year")
	year, err := utils.ValidateYearQueryParam(yearVal)
	if err != nil {
		slog.Error("error converting year value to integer", "error", err)
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	monthVal := r.URL.Query().Get("month")
	month, err := utils.ValidateMonthQueryParam(monthVal)
	if err != nil {
		slog.Error("error converting month value to integer", "error", err)
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	monthlyContributionSummary, err := h.contributionService.ListMonthlyContributionSummary(ctx, year, month, userId)
	if err != nil {
		slog.Error("error fetching contribution type summary for month", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "contribution type overview for month fetched successfully", monthlyContributionSummary)
}

func (h *handler) ListAllContributionTypes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contributionTypes, err := h.contributionService.ListAllContributionTypes(ctx)
	if err != nil {
		slog.Error("error fetching all contribution types", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "all contribution types fetched successfully", contributionTypes)
}

func (h *handler) ConfigureContributionTypeScore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// isAdminValue := ctx.Value(middleware.IsAdminKey)
	// isAdmin, ok := isAdminValue.(bool)
	// if !ok {
	// 	slog.Error("error verifying admin from context")
	// 	status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
	// 	response.WriteJson(w, status, errorMessage, nil)
	// 	return
	// }

	var configureContributionTypeScores []ConfigureContributionTypeScore
	err := json.NewDecoder(r.Body).Decode(&configureContributionTypeScores)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	contributionTypeScores, err := h.contributionService.ConfigureContributionTypeScore(ctx, configureContributionTypeScores)
	if err != nil {
		slog.Error("error configuring contribution type scores", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "contribution types fscores configured successfully", contributionTypeScores)
}
