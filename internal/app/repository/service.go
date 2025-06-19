package repository

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	repositoryRepository repository.RepositoryRepository
	appCfg               config.AppConfig
	httpClient           *http.Client
}

type Service interface {
	GetRepoByRepoId(ctx context.Context, githubRepoId int) (Repository, error)
	FetchRepositoryDetails(ctx context.Context, getUserRepoDetailsUrl string) (FetchRepositoryDetailsResponse, error)
	CreateRepository(ctx context.Context, repoGithubId int, repo FetchRepositoryDetailsResponse) (Repository, error)
}

func NewService(repositoryRepository repository.RepositoryRepository, appCfg config.AppConfig, httpClient *http.Client) Service {
	return &service{
		repositoryRepository: repositoryRepository,
		appCfg:               appCfg,
		httpClient:           httpClient,
	}
}

func (s *service) GetRepoByRepoId(ctx context.Context, repoGithubId int) (Repository, error) {
	repoDetails, err := s.repositoryRepository.GetRepoByGithubId(ctx, nil, repoGithubId)
	if err != nil {
		slog.Error("failed to get repository by github id")
		return Repository{}, err
	}

	return Repository(repoDetails), nil
}

func (s *service) FetchRepositoryDetails(ctx context.Context, getUserRepoDetailsUrl string) (FetchRepositoryDetailsResponse, error) {
	req, err := http.NewRequest("GET", getUserRepoDetailsUrl, nil)
	if err != nil {
		slog.Error("error fetching user repositories details", "error", err)
		return FetchRepositoryDetailsResponse{}, err
	}

	req.Header.Add("Authorization", s.appCfg.GithubPersonalAccessToken)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		slog.Error("error fetching user repositories details", "error", err)
		return FetchRepositoryDetailsResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error freading body", "error", err)
		return FetchRepositoryDetailsResponse{}, err
	}

	var repoDetails FetchRepositoryDetailsResponse
	err = json.Unmarshal(body, &repoDetails)
	if err != nil {
		slog.Error("error unmarshalling fetch repository details body", "error", err)
		return FetchRepositoryDetailsResponse{}, err
	}

	return repoDetails, nil
}

func (s *service) CreateRepository(ctx context.Context, repoGithubId int, repo FetchRepositoryDetailsResponse) (Repository, error) {
	createRepo := Repository{
		GithubRepoId: repoGithubId,
		RepoName:     repo.Name,
		RepoUrl:      repo.RepoUrl,
		Description:  repo.Description,
		LanguagesUrl: repo.LanguagesURL,
		OwnerName:    repo.RepoOwnerName.Login,
		UpdateDate:   repo.UpdateDate,
	}
	repositoryCreated, err := s.repositoryRepository.CreateRepository(ctx, nil, repository.Repository(createRepo))
	if err != nil {
		slog.Error("failed to create repository", "error", err)
		return Repository{}, err
	}

	return Repository(repositoryCreated), nil
}
