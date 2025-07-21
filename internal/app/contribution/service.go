package contribution

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/bigquery"
	repoService "github.com/joshsoftware/code-curiosity-2025/internal/app/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
	"google.golang.org/api/iterator"
)

type service struct {
	bigqueryService        bigquery.Service
	contributionRepository repository.ContributionRepository
	repositoryService      repoService.Service
	userService            user.Service
	httpClient             *http.Client
}

type Service interface {
	ProcessFetchedContributions(ctx context.Context) error
	CreateContribution(ctx context.Context, contributionType string, contributionDetails ContributionResponse, repositoryId int, userId int) (Contribution, error)
	GetContributionScoreDetailsByContributionType(ctx context.Context, contributionType string) (ContributionScore, error)
	FetchUsersAllContributions(ctx context.Context) ([]Contribution, error)
}

func NewService(bigqueryService bigquery.Service, contributionRepository repository.ContributionRepository, repositoryService repoService.Service, userService user.Service, httpClient *http.Client) Service {
	return &service{
		bigqueryService:        bigqueryService,
		contributionRepository: contributionRepository,
		repositoryService:      repositoryService,
		userService:            userService,
		httpClient:             httpClient,
	}
}

func (s *service) ProcessFetchedContributions(ctx context.Context) error {
	contributions, err := s.bigqueryService.FetchDailyContributions(ctx)
	if err != nil {
		slog.Error("error fetching daily contributions", "error", err)
		return err
	}

	for {
		var contribution ContributionResponse
		if err := contributions.Next(&contribution); err == iterator.Done {
			break
		} else if err != nil {
			slog.Error("error iterating contribution rows", "error", err)
			break
		}

		contributionType, err := s.GetContributionType(ctx, contribution)
		if err != nil {
			slog.Error("error getting contribution type", "error", err)
			return err
		}

		var repositoryId int
		repoFetched, err := s.repositoryService.GetRepoByRepoId(ctx, contribution.RepoID) //err no rows
		if err != nil {
			repo, err := s.repositoryService.FetchRepositoryDetails(ctx, contribution.RepoUrl)
			if err != nil {
				slog.Error("error fetching repository details", "error", err)
				return err
			}

			repositoryCreated, err := s.repositoryService.CreateRepository(ctx, contribution.RepoID, repo)
			if err != nil {
				slog.Error("error creating repository", "error", err)
				return err
			}

			repositoryId = repositoryCreated.Id
		} else {
			repositoryId = repoFetched.Id
		}

		user, err := s.userService.GetUserByGithubId(ctx, contribution.ActorID)
		if err != nil {
			slog.Error("error getting user id", "error", err)
			return err
		}

		_, err = s.CreateContribution(ctx, contributionType, contribution, repositoryId, user.Id)
		if err != nil {
			slog.Error("error creating contribution", "error", err)
			return err
		}
	}

	return nil
}

func (s *service) GetContributionType(ctx context.Context, contribution ContributionResponse) (string, error) {
	var contributionPayload map[string]interface{}
	err := json.Unmarshal([]byte(contribution.Payload), &contributionPayload)
	if err != nil {
		slog.Warn("invalid payload", "error", err)
		return "", err
	}

	var action string
	if actionVal, ok := contributionPayload["action"]; ok {
		action = actionVal.(string)
	}

	var pullRequest map[string]interface{}
	var isMerged bool
	if pullRequestPayload, ok := contributionPayload["pull_request"]; ok {
		pullRequest = pullRequestPayload.(map[string]interface{})
		isMerged = pullRequest["merged"].(bool)
	}

	var issue map[string]interface{}
	var stateReason string
	if issuePayload, ok := contributionPayload["issue"]; ok {
		issue = issuePayload.(map[string]interface{})
		stateReason = issue["state_reason"].(string)
	}

	var contributionType string
	switch contribution.Type {
	case "PullRequestEvent":
		if action == "closed" && isMerged {
			contributionType = "PullRequestMerged"
		} else if action == "opened" {
			contributionType = "PullRequestOpened"
		}

	case "IssuesEvent":
		if action == "opened" {
			contributionType = "IssueOpened"
		} else if action == "closed" && stateReason == "not_planned" {
			contributionType = "IssueClosed"
		} else if action == "closed" && stateReason == "completed" {
			contributionType = "IssueResolved"
		}

	case "PushEvent":
		contributionType = "PullRequestUpdated"

	case "IssueCommentEvent":
		contributionType = "IssueComment"

	case "PullRequestComment ":
		contributionType = "PullRequestComment"
	}

	return contributionType, nil
}

func (s *service) CreateContribution(ctx context.Context, contributionType string, contributionDetails ContributionResponse, repositoryId int, userId int) (Contribution, error) {

	contribution := Contribution{
		UserId:           userId,
		RepositoryId:     repositoryId,
		ContributionType: contributionType,
		ContributedAt:    contributionDetails.CreatedAt,
	}

	contributionScoreDetails, err := s.GetContributionScoreDetailsByContributionType(ctx, contributionType)
	if err != nil {
		slog.Error("error occured while getting contribution score details", "error", err)
		return Contribution{}, err
	}

	contribution.ContributionScoreId = contributionScoreDetails.Id
	contribution.BalanceChange = contributionScoreDetails.Score

	contributionResponse, err := s.contributionRepository.CreateContribution(ctx, nil, repository.Contribution(contribution))
	if err != nil {
		slog.Error("error creating contribution", "error", err)
		return Contribution{}, err
	}

	return Contribution(contributionResponse), nil
}

func (s *service) GetContributionScoreDetailsByContributionType(ctx context.Context, contributionType string) (ContributionScore, error) {
	contributionScoreDetails, err := s.contributionRepository.GetContributionScoreDetailsByContributionType(ctx, nil, contributionType)
	if err != nil {
		slog.Error("error occured while getting contribution score details", "error", err)
		return ContributionScore{}, err
	}

	return ContributionScore(contributionScoreDetails), nil
}

func (s *service) FetchUsersAllContributions(ctx context.Context) ([]Contribution, error) {
	usersAllContributions, err := s.contributionRepository.FetchUsersAllContributions(ctx, nil)
	if err != nil {
		slog.Error("error occured while fetching all contributions for user", "error", err)
		return nil, err
	}

	serviceContributions := make([]Contribution, len(usersAllContributions))
	for i, c := range usersAllContributions {
		serviceContributions[i] = Contribution((c))
	}

	return serviceContributions, nil
}
