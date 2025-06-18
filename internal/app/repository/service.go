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
}

type Service interface {
	GetRepoByRepoId(ctx context.Context, githubRepoId int) (Repository, error)
	FetchRepositoryDetails(ctx context.Context, client *http.Client, getUserRepoDetailsUrl string) (FetchRepositoryDetailsResponse, error)
	CreateRepository(ctx context.Context, repoGithubId int, repo FetchRepositoryDetailsResponse) (Repository, error)
	FetchRepositoryLanguages(ctx context.Context, client *http.Client, getRepoLanguagesURL string) (RepoLanguages, error)
	FetchUsersContributedRepos(ctx context.Context, client *http.Client) ([]FetchUsersContributedReposResponse, error)
}

func NewService(repositoryRepository repository.RepositoryRepository, appCfg config.AppConfig) Service {
	return &service{
		repositoryRepository: repositoryRepository,
		appCfg:               appCfg,
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

func (s *service) FetchRepositoryDetails(ctx context.Context, client *http.Client, getUserRepoDetailsUrl string) (FetchRepositoryDetailsResponse, error) {
	req, err := http.NewRequest("GET", getUserRepoDetailsUrl, nil)
	if err != nil {
		slog.Error("error fetching user repositories details", "error", err)
		return FetchRepositoryDetailsResponse{}, err
	}

	req.Header.Add("Authorization", s.appCfg.GithubPersonalAccessToken)

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("error fetching user repositories details", "error", err)
		return FetchRepositoryDetailsResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading body", "error", err)
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

func (s *service) FetchRepositoryLanguages(ctx context.Context, client *http.Client, getRepoLanguagesURL string) (RepoLanguages, error) {
	req, err := http.NewRequest("GET", getRepoLanguagesURL, nil)
	if err != nil {
		slog.Error("error fetching languages for repository", "error", err)
		return RepoLanguages{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("error fetching languages for repository", "error", err)
		return RepoLanguages{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading body", "error", err)
		return RepoLanguages{}, err
	}

	var repoLanguages RepoLanguages
	err = json.Unmarshal(body, &repoLanguages)
	if err != nil {
		slog.Error("error unmarshalling fetch repository languages body", "error", err)
		return RepoLanguages{}, err
	}

	return repoLanguages, nil
}

func (s *service) FetchUsersContributedRepos(ctx context.Context, client *http.Client) ([]FetchUsersContributedReposResponse, error) {
	usersContributedRepos, err := s.repositoryRepository.FetchUsersContributedRepos(ctx, nil)
	if err != nil {
		slog.Error("error fetching users conributed repos")
		return nil, err
	}

	fetchUsersContributedReposResponse := make([]FetchUsersContributedReposResponse, len(usersContributedRepos))

	for i, usersContributedRepo := range usersContributedRepos {
		fetchUsersContributedReposResponse[i].Repository = Repository(usersContributedRepo)

		contributedRepoLanguages, err := s.FetchRepositoryLanguages(ctx, client, usersContributedRepo.LanguagesUrl)
		if err != nil {
			slog.Error("error fetching languages for repository", "error", err)
			return nil, err
		}

		for language := range contributedRepoLanguages {
			fetchUsersContributedReposResponse[i].Languages = append(fetchUsersContributedReposResponse[i].Languages, language)
		}

		userRepoTotalCoins, err := s.repositoryRepository.GetUserRepoTotalCoins(ctx, nil, usersContributedRepo.Id)
		if err != nil {
			slog.Error("error calculating total coins earned by user for the repository")
			return nil, err
		}

		fetchUsersContributedReposResponse[i].TotalCoinsEarned = userRepoTotalCoins
	}

	return fetchUsersContributedReposResponse, nil
}
