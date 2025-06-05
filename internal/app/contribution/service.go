package contribution

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/bigquery"
	repoService "github.com/joshsoftware/code-curiosity-2025/internal/app/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
	"google.golang.org/api/iterator"
)

type service struct {
	bigqueryService        bigquery.Service
	contributionRepository repository.ContributionRepository
	repositoryService      repoService.Service
}

type Service interface {
	ProcessFetchedContributions(ctx context.Context) error
	CreateContribution(ctx context.Context, contributionType string, contributionDetails ContributionResponse, repositoryId int) (Contribution, error)
}

func NewService(bigqueryService bigquery.Service, contributionRepository repository.ContributionRepository, repositoryService repoService.Service) Service {
	return &service{
		bigqueryService:        bigqueryService,
		contributionRepository: contributionRepository,
		repositoryService:      repositoryService,
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

		var contributionPayload map[string]interface{}
		err := json.Unmarshal([]byte(contribution.Payload), &contributionPayload)
		if err != nil {
			slog.Warn("invalid payload", "error", err)
			continue
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

		repoFetched, err := s.repositoryService.GetRepoByGithubId(ctx, contribution.RepoID)
		repositoryId := repoFetched.Id
		if err != nil {
			repo, err := s.repositoryService.FetchRepositoryDetails(ctx, contribution.RepoUrl)
			if err != nil {
				slog.Error("error fetching repository details")
				return err
			}

			repositoryCreated, err := s.repositoryService.CreateRepository(ctx, contribution.RepoID, repo)
			if err != nil {
				slog.Error("error creating repository", "error", err)
				return err
			}

			repositoryId = repositoryCreated.Id
		}

		_, err = s.CreateContribution(ctx, contributionType, contribution, repositoryId)
		if err != nil {
			slog.Error("error creating contribution", "error", err)
			return err
		}
	}
	return nil
}

func (s *service) CreateContribution(ctx context.Context, contributionType string, contributionDetails ContributionResponse, repositoryId int) (Contribution, error) {
	contribution := Contribution{
		UserId:           contributionDetails.ActorID,
		RepositoryId:     repositoryId,
		ContributionType: contributionType,
		//get id and balance from contribution_score_id table by sending it contribution_type (hardcoded for now)
		ContributionScoreId: 1,
		BalanceChange:       10,
		ContributedAt:       contributionDetails.CreatedAt,
	}

	contributionResponse, err := s.contributionRepository.CreateContribution(ctx, nil, repository.Contribution(contribution))
	if err != nil {
		slog.Error("error creating contribution", "error", err)
		return Contribution{}, err
	}

	return Contribution(contributionResponse), nil
}
