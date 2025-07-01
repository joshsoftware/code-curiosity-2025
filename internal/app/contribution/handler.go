package contribution

import (
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	contributionService Service
}

type Handler interface {
	FetchUserLatestContributions(w http.ResponseWriter, r *http.Request)
	FetchUserContributions(w http.ResponseWriter, r *http.Request)
}

func NewHandler(contributionService Service) Handler {
	return &handler{
		contributionService: contributionService,
	}
}

func (h *handler) FetchUserLatestContributions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.contributionService.ProcessFetchedContributions(ctx)
	if err != nil {
		slog.Error("error fetching latest contributions", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "contribution fetched successfully", nil)
}

func (h *handler) FetchUserContributions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userContributions, err := h.contributionService.FetchUserContributions(ctx)
	if err != nil {
		slog.Error("error fetching user contributions")
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "user contributions fetched successfully", userContributions)
}
