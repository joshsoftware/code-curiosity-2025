package repository

import (
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	repositoryService Service
}

type Handler interface {
	FetchUsersContributedRepos(w http.ResponseWriter, r *http.Request)
}

func NewHandler(repositoryService Service) Handler {
	return &handler{
		repositoryService: repositoryService,
	}
}

func (h *handler) FetchUsersContributedRepos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	client := &http.Client{}

	usersContributedRepos, err := h.repositoryService.FetchUsersContributedRepos(ctx, client)
	if err != nil {
		slog.Error("error fetching users conributed repos")
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "users contributed repositories fetched successfully", usersContributedRepos)
}
