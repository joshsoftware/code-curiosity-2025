package repository

import (
	"context"
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
	FetchUsersAllContributions(ctx context.Context, tx *sqlx.Tx) ([]Contribution, error)
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
	contributed_at
	)
	VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING *`

	getContributionScoreDetailsByContributionTypeQuery = `SELECT * from contribution_score where contribution_type=$1`

	fetchUsersAllContributionsQuery = `SELECT * from contributions where user_id=$1 order by contributed_at desc`
)

func (cr *contributionRepository) CreateContribution(ctx context.Context, tx *sqlx.Tx, contributionInfo Contribution) (Contribution, error) {
	executer := cr.BaseRepository.initiateQueryExecuter(tx)

	var contribution Contribution
	err := executer.QueryRowContext(ctx, createContributionQuery,
		contributionInfo.UserId,
		contributionInfo.RepositoryId,
		contributionInfo.ContributionScoreId,
		contributionInfo.ContributionType,
		contributionInfo.BalanceChange,
		contributionInfo.ContributedAt,
	).Scan(
		&contribution.Id,
		&contribution.UserId,
		&contribution.RepositoryId,
		&contribution.ContributionScoreId,
		&contribution.ContributionType,
		&contribution.BalanceChange,
		&contribution.ContributedAt,
		&contribution.CreatedAt,
		&contribution.UpdatedAt,
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
	err := executer.QueryRowContext(ctx, getContributionScoreDetailsByContributionTypeQuery, contributionType).Scan(
		&contributionScoreDetails.Id,
		&contributionScoreDetails.AdminId,
		&contributionScoreDetails.ContributionType,
		&contributionScoreDetails.Score,
		&contributionScoreDetails.CreatedAt,
		&contributionScoreDetails.UpdatedAt,
	)
	if err != nil {
		slog.Error("error occured while getting contribution score details", "error", err)
		return ContributionScore{}, err
	}

	return contributionScoreDetails, nil
}

func (cr *contributionRepository) FetchUsersAllContributions(ctx context.Context, tx *sqlx.Tx) ([]Contribution, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return nil, apperrors.ErrInternalServer
	}

	executer := cr.BaseRepository.initiateQueryExecuter(tx)

	rows, err := executer.QueryContext(ctx, fetchUsersAllContributionsQuery, userId)
	if err != nil {
		slog.Error("error fetching all contributions for user", "error", err)
		return nil, apperrors.ErrFetchingAllContributions
	}
	defer rows.Close()

	var usersAllContributions []Contribution
	for rows.Next() {
		var currentContribution Contribution
		if err = rows.Scan(
			&currentContribution.Id,
			&currentContribution.UserId,
			&currentContribution.RepositoryId,
			&currentContribution.ContributionScoreId,
			&currentContribution.ContributionType,
			&currentContribution.BalanceChange,
			&currentContribution.ContributedAt,
			&currentContribution.CreatedAt, &currentContribution.UpdatedAt); err != nil {
			return nil, err
		}

		usersAllContributions = append(usersAllContributions, currentContribution)
	}

	return usersAllContributions, nil
}
