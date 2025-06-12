package contribution

import (
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	contributionService Service
}

type Handler interface {
	FetchUserLatestContributions(w http.ResponseWriter, r *http.Request)
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
		slog.Error("error fetching latest contributions")
		return
	}

	response.WriteJson(w, http.StatusOK, "contribution fetched successfully", nil)
}
