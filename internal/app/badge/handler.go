package badge

import (
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type badgeHandler struct {
	badgeService BadgeService
}

type BadgeHandler interface {
	GetBadgeDetailsOfUser(w http.ResponseWriter, r *http.Request)
}

func NewBadgeHandler(badgeService BadgeService) BadgeHandler {
	return &badgeHandler{
		badgeService: badgeService,
	}
}

func (bh *badgeHandler) GetBadgeDetailsOfUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	badges, err := bh.badgeService.GetBadgeDetailsOfUser(ctx)

	if err != nil {
		slog.Error("(handler) Failed to get badge details of user", "Error", err)
		status, errorMsg := apperrors.MapError(err)
		response.WriteJson(w, status, errorMsg, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "Badges fetched successfully", badges)
}
