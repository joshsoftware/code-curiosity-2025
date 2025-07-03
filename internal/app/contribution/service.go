package contribution

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/bigquery"
	repoService "github.com/joshsoftware/code-curiosity-2025/internal/app/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/transaction"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
	"google.golang.org/api/iterator"
)

// github event names
const (
	pullRequestEvent  = "PullRequestEvent"
	issuesEvent       = "IssuesEvent"
	pushEvent         = "PushEvent"
	issueCommentEvent = "IssueCommentEvent"
)

// app contribution types
const (
	pullRequestMerged  = "PullRequestMerged"
	pullRequestOpened  = "PullRequestOpened"
	issueOpened        = "IssueOpened"
	issueClosed        = "IssueClosed"
	issueResolved      = "IssueResolved"
	pullRequestUpdated = "PullRequestUpdated"
	issueComment       = "IssueComment"
	pullRequestComment = "PullRequestComment"
)

// payload
const (
	payloadActionKey      = "action"
	payloadPullRequestKey = "pull_request"
	PayloadMergedKey      = "merged"
	PayloadIssueKey       = "issue"
	PayloadStateReasonKey = "state_reason"
	PayloadClosedKey      = "closed"
	PayloadOpenedKey      = "opened"
	PayloadNotPlannedKey  = "not_planned"
	PayloadCompletedKey   = "completed"
)

type service struct {
	bigqueryService        bigquery.Service
	contributionRepository repository.ContributionRepository
	repositoryService      repoService.Service
	userService            user.Service
	transactionService     transaction.Service
	httpClient             *http.Client
}

type Service interface {
	ProcessFetchedContributions(ctx context.Context) error
	ProcessEachContribution(ctx context.Context, contribution ContributionResponse) error
	CreateContribution(ctx context.Context, contributionType string, contributionDetails ContributionResponse, repositoryId int, userId int) (Contribution, error)
	GetContributionScoreDetailsByContributionType(ctx context.Context, contributionType string) (ContributionScore, error)
	FetchUserContributions(ctx context.Context) ([]Contribution, error)
	GetContributionByGithubEventId(ctx context.Context, githubEventId string) (Contribution, error)
	CreateContributionTransaction(ctx context.Context, userId int, contributionDetails Contribution) (Transaction, error)
}

func NewService(bigqueryService bigquery.Service, contributionRepository repository.ContributionRepository, repositoryService repoService.Service, userService user.Service, transactionService transaction.Service, httpClient *http.Client) Service {
	return &service{
		bigqueryService:        bigqueryService,
		contributionRepository: contributionRepository,
		repositoryService:      repositoryService,
		userService:            userService,
		transactionService:     transactionService,
		httpClient:             httpClient,
	}
}

func (s *service) ProcessFetchedContributions(ctx context.Context) error {
	contributions, err := s.bigqueryService.FetchDailyContributions(ctx)
	if err != nil {
		slog.Error("error fetching daily contributions", "error", err)
		return apperrors.ErrFetchingFromBigquery
	}

	//using a local copy here to copy contribution so that I can implement retry mechanism in future
	//thinking of batch processing to be implemented later on, to handle memory overflow
	var fetchedContributions []ContributionResponse

	for {
		var contribution ContributionResponse
		err := contributions.Next(&contribution)
		if err != nil {
			if err == iterator.Done {
				break
			}

			slog.Error("error iterating contribution rows", "error", err)
			return apperrors.ErrNextContribution
		}

		fetchedContributions = append(fetchedContributions, contribution)
	}

	for _, contribution := range fetchedContributions {
		err := s.ProcessEachContribution(ctx, contribution)
		if err != nil {
			slog.Error("error processing contribution with github event id", "github event id", "error", contribution.ID, err)
			return err
		}
	}

	return nil
}

func (s *service) ProcessEachContribution(ctx context.Context, contribution ContributionResponse) error {
	_, err := s.GetContributionByGithubEventId(ctx, contribution.ID)
	if err == nil {
		return nil
	}

	if err != apperrors.ErrContributionNotFound {
		slog.Error("error fetching contribution by github event id", "error", err)
		return err
	}

	var repositoryId int
	repoFetched, err := s.repositoryService.GetRepoByGithubId(ctx, contribution.RepoID)
	if err == nil {
		repositoryId = repoFetched.Id
	} else if err == apperrors.ErrRepoNotFound {
		repositoryCreated, err := s.repositoryService.CreateRepository(ctx, contribution.RepoID, contribution.RepoUrl)
		if err != nil {
			slog.Error("error creating repository", "error", err)
			return err
		}

		repositoryId = repositoryCreated.Id
	} else {
		slog.Error("error fetching repo by repo id", "error", err)
		return err
	}

	user, err := s.userService.GetUserByGithubId(ctx, contribution.ActorID)
	if err != nil {
		slog.Error("error getting user id", "error", err)
		return err
	}

	contributionType, err := s.GetContributionType(ctx, contribution)
	if err != nil {
		slog.Error("error getting contribution type", "error", err)
		return err
	}

	createdContribution, err := s.CreateContribution(ctx, contributionType, contribution, repositoryId, user.Id)
	if err != nil {
		slog.Error("error creating contribution", "error", err)
		return err
	}

	_, err = s.CreateContributionTransaction(ctx, user.Id, createdContribution)
	if err != nil {
		slog.Error("error creating transaction for current contribution", "error", err)
		return err
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
	if actionVal, ok := contributionPayload[payloadActionKey]; ok {
		action = actionVal.(string)
	}

	var pullRequest map[string]interface{}
	var isMerged bool
	if pullRequestPayload, ok := contributionPayload[payloadPullRequestKey]; ok {
		pullRequest = pullRequestPayload.(map[string]interface{})
		isMerged = pullRequest[PayloadMergedKey].(bool)
	}

	var issue map[string]interface{}
	var stateReason string
	if issuePayload, ok := contributionPayload[PayloadIssueKey]; ok {
		issue = issuePayload.(map[string]interface{})
		stateReason = issue[PayloadStateReasonKey].(string)
	}

	var contributionType string
	switch contribution.Type {
	case pullRequestEvent:
		if action == PayloadClosedKey && isMerged {
			contributionType = pullRequestMerged
		} else if action == PayloadOpenedKey {
			contributionType = pullRequestOpened
		}

	case issuesEvent:
		if action == PayloadOpenedKey {
			contributionType = issueOpened
		} else if action == PayloadClosedKey && stateReason == PayloadNotPlannedKey {
			contributionType = issueClosed
		} else if action == PayloadClosedKey && stateReason == PayloadCompletedKey {
			contributionType = issueResolved
		}

	case pushEvent:
		contributionType = pullRequestUpdated

	case issueCommentEvent:
		contributionType = issueComment

	case pullRequestComment:
		contributionType = pullRequestComment
	}

	return contributionType, nil
}

func (s *service) CreateContribution(ctx context.Context, contributionType string, contributionDetails ContributionResponse, repositoryId int, userId int) (Contribution, error) {

	contribution := Contribution{
		UserId:           userId,
		RepositoryId:     repositoryId,
		ContributionType: contributionType,
		ContributedAt:    contributionDetails.CreatedAt,
		GithubEventId:    contributionDetails.ID,
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

func (s *service) FetchUserContributions(ctx context.Context) ([]Contribution, error) {
	userContributions, err := s.contributionRepository.FetchUserContributions(ctx, nil)
	if err != nil {
		slog.Error("error occured while fetching user contributions", "error", err)
		return nil, err
	}

	serviceContributions := make([]Contribution, len(userContributions))
	for i, c := range userContributions {
		serviceContributions[i] = Contribution((c))
	}

	return serviceContributions, nil
}

func (s *service) GetContributionByGithubEventId(ctx context.Context, githubEventId string) (Contribution, error) {
	contribution, err := s.contributionRepository.GetContributionByGithubEventId(ctx, nil, githubEventId)
	if err != nil {
		slog.Error("error fetching contribution by github event id", "error", err)
		return Contribution{}, err
	}

	return Contribution(contribution), nil
}

func (s *service) CreateContributionTransaction(ctx context.Context, userId int, contributionDetails Contribution) (Transaction, error) {
	transactionInfo := Transaction{
		UserId:            userId,
		ContributionId:    contributionDetails.Id,
		IsRedeemed:        false,
		IsGained:          true,
		TransactedBalance: contributionDetails.BalanceChange,
		TransactedAt:      contributionDetails.ContributedAt,
	}
	transaction, err := s.transactionService.CreateTransaction(ctx, transaction.Transaction(transactionInfo))
	if err != nil {
		slog.Error("error creating transaction for current contribution", "error", err)
		return Transaction{}, err
	}

	return Transaction(transaction), nil
}
