package badge

import (
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	badgeService Service
}

type Handler interface {
	GetBadgeDetailsOfUser(w http.ResponseWriter, r *http.Request)
}

func NewHandler(badgeService Service) Handler {
	return &handler{
		badgeService: badgeService,
	}
}

func (h *handler) GetBadgeDetailsOfUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdValue := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMsg := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMsg, nil)
		return
	}

	badges, err := h.badgeService.GetBadgeDetailsOfUser(ctx, userId)

	if err != nil {
		slog.Error("failed to get badge details of user", "Error", err)
		status, errorMsg := apperrors.MapError(err)
		response.WriteJson(w, status, errorMsg, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "badges fetched successfully", badges)
}
