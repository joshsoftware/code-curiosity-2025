package repository

import (
	"context"
	"log/slog"
	"math"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/github"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	repositoryRepository repository.RepositoryRepository
	githubService        github.Service
}

type Service interface {
	GetRepoByGithubId(ctx context.Context, githubRepoId int) (Repository, error)
	GetRepoByRepoId(ctx context.Context, repoId int) (Repository, error)
	CreateRepository(ctx context.Context, repoGithubId int, ContributionRepoDetailsUrl string) (Repository, error)
	HandleRepositoryCreation(ctx context.Context, contribution ContributionResponse) (Repository, error)
	FetchUsersContributedRepos(ctx context.Context, client *http.Client) ([]FetchUsersContributedReposResponse, error)
	FetchUserContributionsInRepo(ctx context.Context, githubRepoId int) ([]Contribution, error)
	CalculateLanguagePercentInRepo(ctx context.Context, repoLanguages RepoLanguages) ([]LanguagePercent, error)
}

func NewService(repositoryRepository repository.RepositoryRepository, githubService github.Service) Service {
	return &service{
		repositoryRepository: repositoryRepository,
		githubService:        githubService,
	}
}

func (s *service) GetRepoByGithubId(ctx context.Context, repoGithubId int) (Repository, error) {
	repoDetails, err := s.repositoryRepository.GetRepoByGithubId(ctx, nil, repoGithubId)
	if err != nil {
		slog.Error("failed to get repository by repo github id", "error", err)
		return Repository{}, err
	}

	return Repository(repoDetails), nil
}

func (s *service) GetRepoByRepoId(ctx context.Context, repobId int) (Repository, error) {
	repoDetails, err := s.repositoryRepository.GetRepoByRepoId(ctx, nil, repobId)
	if err != nil {
		slog.Error("failed to get repository by repo id", "error", err)
		return Repository{}, err
	}

	return Repository(repoDetails), nil
}

func (s *service) CreateRepository(ctx context.Context, repoGithubId int, ContributionRepoDetailsUrl string) (Repository, error) {
	repo, err := s.githubService.FetchRepositoryDetails(ctx, ContributionRepoDetailsUrl)
	if err != nil {
		slog.Error("error fetching user repositories details", "error", err)
		return Repository{}, err
	}

	createRepo := Repository{
		GithubRepoId:    repoGithubId,
		RepoName:        repo.Name,
		RepoUrl:         repo.RepoUrl,
		Description:     repo.Description,
		LanguagesUrl:    repo.LanguagesURL,
		OwnerName:       repo.RepoOwnerName.Login,
		UpdateDate:      repo.UpdateDate,
		ContributorsUrl: repo.ContributorsUrl,
	}
	repositoryCreated, err := s.repositoryRepository.CreateRepository(ctx, nil, repository.Repository(createRepo))
	if err != nil {
		slog.Error("failed to create repository", "error", err)
		return Repository{}, err
	}

	return Repository(repositoryCreated), nil
}

func (s *service) HandleRepositoryCreation(ctx context.Context, contribution ContributionResponse) (Repository, error) {
	obtainedRepository, err := s.GetRepoByGithubId(ctx, contribution.RepoID)
	if err != nil {
		if err == apperrors.ErrRepoNotFound {
			obtainedRepository, err = s.CreateRepository(ctx, contribution.RepoID, contribution.RepoUrl)
			if err != nil {
				slog.Error("error creating repository", "error", err)
				return Repository{}, err
			}
		} else {
			slog.Error("error fetching repo by repo id", "error", err)
			return Repository{}, err
		}
	}

	return obtainedRepository, nil
}

func (s *service) FetchUsersContributedRepos(ctx context.Context, client *http.Client) ([]FetchUsersContributedReposResponse, error) {
	usersContributedRepos, err := s.repositoryRepository.FetchUsersContributedRepos(ctx, nil)
	if err != nil {
		slog.Error("error fetching users conributed repos", "error", err)
		return nil, err
	}

	fetchUsersContributedReposResponse := make([]FetchUsersContributedReposResponse, len(usersContributedRepos))

	for i, usersContributedRepo := range usersContributedRepos {
		fetchUsersContributedReposResponse[i].Repository = Repository(usersContributedRepo)

		contributedRepoLanguages, err := s.githubService.FetchRepositoryLanguages(ctx, usersContributedRepo.LanguagesUrl)
		if err != nil {
			slog.Error("error fetching languages for repository", "error", err)
			return nil, err
		}

		for language := range contributedRepoLanguages {
			fetchUsersContributedReposResponse[i].Languages = append(fetchUsersContributedReposResponse[i].Languages, language)
		}

		userRepoTotalCoins, err := s.repositoryRepository.GetUserRepoTotalCoins(ctx, nil, usersContributedRepo.Id)
		if err != nil {
			slog.Error("error calculating total coins earned by user for the repository", "error", err)
			return nil, err
		}

		fetchUsersContributedReposResponse[i].TotalCoinsEarned = userRepoTotalCoins
	}

	return fetchUsersContributedReposResponse, nil
}

func (s *service) FetchUserContributionsInRepo(ctx context.Context, githubRepoId int) ([]Contribution, error) {
	userContributionsInRepo, err := s.repositoryRepository.FetchUserContributionsInRepo(ctx, nil, githubRepoId)
	if err != nil {
		slog.Error("error fetching users contribution in repository", "error", err)
		return nil, err
	}

	serviceUserContributionsInRepo := make([]Contribution, len(userContributionsInRepo))
	for i, c := range userContributionsInRepo {
		serviceUserContributionsInRepo[i] = Contribution(c)
	}

	return serviceUserContributionsInRepo, nil
}

func (s *service) CalculateLanguagePercentInRepo(ctx context.Context, repoLanguages RepoLanguages) ([]LanguagePercent, error) {
	var total int
	for _, bytes := range repoLanguages {
		total += bytes
	}

	var langPercent []LanguagePercent

	for lang, bytes := range repoLanguages {
		percentage := (float64(bytes) / float64(total)) * 100
		langPercent = append(langPercent, LanguagePercent{
			Name:       lang,
			Bytes:      bytes,
			Percentage: math.Round(percentage*10) / 10,
		})
	}

	return langPercent, nil
}
