package github

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
)

type service struct {
	appCfg     config.AppConfig
	httpClient *http.Client
}

type Service interface {
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

func (s *service) FetchRepositoryContributors(ctx context.Context, client *http.Client, getRepoContributorsURl string) ([]FetchRepoContributorsResponse, error) {
	req, err := http.NewRequest("GET", getRepoContributorsURl, nil)
	if err != nil {
		slog.Error("error fetching contributors for repository", "error", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("error fetching contributors for repository", "error", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading body", "error", err)
		return nil, err
	}

	var repoContributors []FetchRepoContributorsResponse
	err = json.Unmarshal(body, &repoContributors)
	if err != nil {
		slog.Error("error unmarshalling fetch contributors body", "error", err)
		return nil, err
	}

	return repoContributors, nil
}
