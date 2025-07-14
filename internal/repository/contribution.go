package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
)

type contributionRepository struct {
	BaseRepository
}

type ContributionRepository interface {
	RepositoryTransaction
	CreateContribution(ctx context.Context, tx *sqlx.Tx, contributionDetails Contribution) (Contribution, error)
	GetContributionScoreDetailsByContributionType(ctx context.Context, tx *sqlx.Tx, contributionType string) (ContributionScore, error)
	FetchUserContributions(ctx context.Context, tx *sqlx.Tx) ([]Contribution, error)
	GetContributionByGithubEventId(ctx context.Context, tx *sqlx.Tx, githubEventId string) (Contribution, error)
}

func NewContributionRepository(db *sqlx.DB) ContributionRepository {
	return &contributionRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	createContributionQuery = `
	INSERT INTO contributions (
	user_id,
	repository_id, 
	contribution_score_id, 
	contribution_type, 
	balance_change, 
	contributed_at,
	github_event_id
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING *`

	getContributionScoreDetailsByContributionTypeQuery = `SELECT * from contribution_score where contribution_type=$1`

	fetchUserContributionsQuery = `SELECT * from contributions where user_id=$1 order by contributed_at desc`

	GetContributionByGithubEventIdQuery = `SELECT * from contributions where github_event_id=$1`
)

func (cr *contributionRepository) CreateContribution(ctx context.Context, tx *sqlx.Tx, contributionInfo Contribution) (Contribution, error) {
	executer := cr.BaseRepository.initiateQueryExecuter(tx)

	var contribution Contribution
	err := executer.GetContext(ctx, &contribution, createContributionQuery,
		contributionInfo.UserId,
		contributionInfo.RepositoryId,
		contributionInfo.ContributionScoreId,
		contributionInfo.ContributionType,
		contributionInfo.BalanceChange,
		contributionInfo.ContributedAt,
		contributionInfo.GithubEventId,
	)
	if err != nil {
		slog.Error("error occured while inserting contributions", "error", err)
		return Contribution{}, apperrors.ErrContributionCreationFailed
	}

	return contribution, err
}

func (cr *contributionRepository) GetContributionScoreDetailsByContributionType(ctx context.Context, tx *sqlx.Tx, contributionType string) (ContributionScore, error) {
	executer := cr.BaseRepository.initiateQueryExecuter(tx)

	var contributionScoreDetails ContributionScore
	err := executer.GetContext(ctx, &contributionScoreDetails, getContributionScoreDetailsByContributionTypeQuery, contributionType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("no contribution score details found for contribution type", "contributionType", contributionType)
			return ContributionScore{}, apperrors.ErrContributionScoreNotFound
		}

		slog.Error("error occured while getting contribution score details", "error", err)
		return ContributionScore{}, err
	}

	return contributionScoreDetails, nil
}

func (cr *contributionRepository) FetchUserContributions(ctx context.Context, tx *sqlx.Tx) ([]Contribution, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return nil, apperrors.ErrInternalServer
	}

	executer := cr.BaseRepository.initiateQueryExecuter(tx)

	var userContributions []Contribution
	err := executer.SelectContext(ctx, &userContributions, fetchUserContributionsQuery, userId)
	if err != nil {
		slog.Error("error fetching user contributions", "error", err)
		return nil, apperrors.ErrFetchingAllContributions
	}

	return userContributions, nil
}

func (cr *contributionRepository) GetContributionByGithubEventId(ctx context.Context, tx *sqlx.Tx, githubEventId string) (Contribution, error) {
	executer := cr.BaseRepository.initiateQueryExecuter(tx)

	var contribution Contribution
	err := executer.GetContext(ctx, &contribution, GetContributionByGithubEventIdQuery, githubEventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("contribution not found", "error", err)
			return Contribution{}, apperrors.ErrContributionNotFound
		}
		slog.Error("error fetching contribution by github event id", "error", err)
		return Contribution{}, apperrors.ErrFetchingContribution
	}

	return contribution, nil

}
