package github

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/utils"
)

type service struct {
	appCfg     config.AppConfig
	httpClient *http.Client
}

type Service interface {
	configureGithubApiHeaders() map[string]string
	FetchRepositoryDetails(ctx context.Context, getUserRepoDetailsUrl string) (FetchRepositoryDetailsResponse, error)
	FetchRepositoryLanguages(ctx context.Context, client *http.Client, getRepoLanguagesURL string) (RepoLanguages, error)
	FetchRepositoryContributors(ctx context.Context, client *http.Client, getRepoContributorsURl string) ([]FetchRepoContributorsResponse, error)
}

func NewService(appCfg config.AppConfig, httpClient *http.Client) Service {
	return &service{
		appCfg:     appCfg,
		httpClient: httpClient,
	}
}

func (s *service) configureGithubApiHeaders() map[string]string {
	return map[string]string{
		AuthorizationKey: s.appCfg.GithubPersonalAccessToken,
	}
}

func (s *service) FetchRepositoryDetails(ctx context.Context, getUserRepoDetailsUrl string) (FetchRepositoryDetailsResponse, error) {
	headers := s.configureGithubApiHeaders()

	body, err := utils.DoGet(s.httpClient, getUserRepoDetailsUrl, headers)
	if err != nil {
		slog.Error("error making a GET request", "error", err)
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

func (s *service) FetchRepositoryLanguages(ctx context.Context, client *http.Client, getRepoLanguagesURL string) (RepoLanguages, error) {
	headers := s.configureGithubApiHeaders()

	body, err := utils.DoGet(s.httpClient, getRepoLanguagesURL, headers)
	if err != nil {
		slog.Error("error making a GET request", "error", err)
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

func (s *service) FetchRepositoryContributors(ctx context.Context, client *http.Client, getRepoContributorsURl string) ([]FetchRepoContributorsResponse, error) {
	headers := s.configureGithubApiHeaders()

	body, err := utils.DoGet(s.httpClient, getRepoContributorsURl, headers)
	if err != nil {
		slog.Error("error making a GET request", "error", err)
		return []FetchRepoContributorsResponse{}, err
	}

	var repoContributors []FetchRepoContributorsResponse
	err = json.Unmarshal(body, &repoContributors)
	if err != nil {
		slog.Error("error unmarshalling fetch contributors body", "error", err)
		return nil, err
	}

	return repoContributors, nil
}
