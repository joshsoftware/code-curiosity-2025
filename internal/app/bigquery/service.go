package bigquery

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	bq "cloud.google.com/go/bigquery"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	bigqueryInstance config.Bigquery
	userRepository   repository.UserRepository
}

type Service interface {
	FetchDailyContributions(ctx context.Context) (*bq.RowIterator, error)
}

func NewService(bigqueryInstance config.Bigquery) Service {
	return &service{
		bigqueryInstance: bigqueryInstance,
	}
}

func (s *service) FetchDailyContributions(ctx context.Context) (*bq.RowIterator, error) {
	YesterdayDate := time.Now().AddDate(0, 0, -1)
	YesterdayYearMonthDay := YesterdayDate.Format("20030101")

	usersNamesList, err := s.userRepository.GetAllUsersGithubUsernames(ctx, nil)
	if err != nil {
		slog.Error("error fetching users github usernames")
		return nil, apperrors.ErrInternalServer
	}

	var quotedUsernamesList []string
	for _, username := range usersNamesList {
		quotedUsernamesList = append(quotedUsernamesList, fmt.Sprintf("'%s'", username))
	}

	githubUsernames := strings.Join(quotedUsernamesList, ",")
	fetchDailyContributionsQuery := fmt.Sprintf(` 
SELECT 
  id,
  type,
  public,
  actor.id AS actor_id,
  actor.login AS actor_login,
  actor.gravatar_id AS actor_gravatar_id,
  actor.url AS actor_url,
  actor.avatar_url AS actor_avatar_url,
  repo.id AS repo_id,
  repo.name AS repo_name,
  repo.url AS repo_url,
  payload,
  created_at,
  other
FROM 
  githubarchive.day.%s
WHERE 
  type IN (
    'IssuesEvent', 
    'PullRequestEvent', 
    'PullRequestReviewEvent', 
    'IssueCommentEvent', 
    'PullRequestReviewCommentEvent'
  )
  AND (
    actor.login IN (%s) OR
    JSON_EXTRACT_SCALAR(payload, "$.pull_request.user.login") IN (%s)
  )
`, YesterdayYearMonthDay, githubUsernames, githubUsernames)

	bigqueryQuery := s.bigqueryInstance.Client.Query(fetchDailyContributionsQuery)
	contributionRows, err := bigqueryQuery.Read(ctx)
	if err != nil {
		slog.Error("error fetching contributions", "error", err)
		return nil, err
	}

	return contributionRows, err
}
