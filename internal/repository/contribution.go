package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type contributionRepository struct {
	BaseRepository
}

type ContributionRepository interface {
	RepositoryTransaction
	CreateContribution(ctx context.Context, tx *sqlx.Tx, contributionDetails Contribution) (Contribution, error)
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
	created_at,
	updated_at
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING *`
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
		time.Now().Unix(),
		time.Now().Unix(),
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
