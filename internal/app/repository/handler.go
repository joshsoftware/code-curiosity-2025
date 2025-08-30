package repository

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	repositoryService Service
}

type Handler interface {
	FetchUsersContributedRepos(w http.ResponseWriter, r *http.Request)
	FetchParticularRepoDetails(w http.ResponseWriter, r *http.Request)
	FetchUserContributionsInRepo(w http.ResponseWriter, r *http.Request)
	FetchLanguagePercentInRepo(w http.ResponseWriter, r *http.Request)
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
		slog.Error("error fetching users conributed repos", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "users contributed repositories fetched successfully", usersContributedRepos)
}

func (h *handler) FetchParticularRepoDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repoIdPath := r.PathValue("repo_id")
	repoId, err := strconv.Atoi(repoIdPath)
	if err != nil {
		slog.Error("error getting repo id from request url", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	repoDetails, err := h.repositoryService.GetRepoByRepoId(ctx, repoId)
	if err != nil {
		slog.Error("error fetching particular repo details", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "repository details fetched successfully", repoDetails)
}

func (h *handler) FetchParticularRepoContributors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	client := &http.Client{}

	repoIdPath := r.PathValue("repo_id")
	repoId, err := strconv.Atoi(repoIdPath)
	if err != nil {
		slog.Error("error getting repo id from request url", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	repoDetails, err := h.repositoryService.GetRepoByRepoId(ctx, repoId)
	if err != nil {
		slog.Error("error fetching particular repo details", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	repoContributors, err := h.repositoryService.FetchRepositoryContributors(ctx, client, repoDetails.ContributorsUrl)
	if err != nil {
		slog.Error("error fetching repo contributors", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "contributors for repo fetched successfully", repoContributors)
}

func (h *handler) FetchUserContributionsInRepo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	repoIdPath := r.PathValue("repo_id")
	repoId, err := strconv.Atoi(repoIdPath)
	if err != nil {
		slog.Error("error getting repo id from request url", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	usersContributionsInRepo, err := h.repositoryService.FetchUserContributionsInRepo(ctx, repoId)
	if err != nil {
		slog.Error("error fetching users contribution in repository", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "users contribution for repository fetched successfully", usersContributionsInRepo)
}

func (h *handler) FetchLanguagePercentInRepo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	client := &http.Client{}

	repoIdPath := r.PathValue("repo_id")
	repoId, err := strconv.Atoi(repoIdPath)
	if err != nil {
		slog.Error("error getting repo id from request url", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	repoDetails, err := h.repositoryService.GetRepoByRepoId(ctx, repoId)
	if err != nil {
		slog.Error("error fetching particular repo details", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	repoLanguages, err := h.repositoryService.FetchRepositoryLanguages(ctx, client, repoDetails.LanguagesUrl)
	if err != nil {
		slog.Error("error fetching particular repo languages", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	langPercent, err := h.repositoryService.CalculateLanguagePercentInRepo(ctx, repoLanguages)
	if err != nil {
		slog.Error("error fetching particular repo languages", "error", err)
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "language percentages for repo fetched successfully", langPercent)
}
