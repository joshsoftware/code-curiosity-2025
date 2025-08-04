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

func NewService(bigqueryInstance config.Bigquery, userRepository repository.UserRepository) Service {
	return &service{
		bigqueryInstance: bigqueryInstance,
		userRepository:   userRepository,
	}
}

func (s *service) FetchDailyContributions(ctx context.Context) (*bq.RowIterator, error) {
	usersNamesList, err := s.userRepository.GetAllUsersGithubUsernames(ctx, nil)
	if err != nil {
		slog.Error("error fetching users github usernames")
		return nil, apperrors.ErrInternalServer
	}

	var quotedUsernamesList []string
	for _, username := range usersNamesList {
		quotedUsernamesList = append(quotedUsernamesList, fmt.Sprintf("'%s'", username))
	}

	YesterdayDate := time.Now().AddDate(0, 0, -1)
	YesterdayYearMonthDay := YesterdayDate.Format("20060102")

	githubUsernames := strings.Join(quotedUsernamesList, ",")
	fetchDailyContributionsQuery := fmt.Sprintf(DailyQuery, YesterdayYearMonthDay, githubUsernames)

	bigqueryQuery := s.bigqueryInstance.Client.Query(fetchDailyContributionsQuery)
	contributionRows, err := bigqueryQuery.Read(ctx)
	if err != nil {
		slog.Error("error fetching contributions", "error", err)
		return nil, err
	}

	return contributionRows, err
}
