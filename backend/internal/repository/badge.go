package repository

import (
	"context"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type badgeRepository struct {
	BaseRepository
}

type BadgeRepository interface {
	RepositoryTransaction
	GetUserCurrentMonthBadge(ctx context.Context, tx *sqlx.Tx, userId int) (Badge, error)
	CreateBadge(ctx context.Context, tx *sqlx.Tx, userId int, badgeType string) (Badge, error)
	GetBadgeDetailsOfUser(ctx context.Context, tx *sqlx.Tx, userId int) ([]Badge, error)
}

func NewBadgeRepository(db *sqlx.DB) BadgeRepository {
	return &badgeRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	createBadgeQuery = `
	INSERT INTO badges(
	user_id,
	badge_type,
	earned_at
	)
	VALUES($1, $2, $3)
	RETURNING *`

	getBadgeDetailsOfUserQuery = "SELECT * FROM badges WHERE user_id = $1 ORDER BY earned_at DESC"

	getUserCurrentMonthBadgeQuery = `
	SELECT * FROM badges
	WHERE user_id = $1
  	AND earned_at >= DATE_TRUNC('month', CURRENT_DATE)
  	AND earned_at < DATE_TRUNC('month', CURRENT_DATE + INTERVAL '1 month')`
)

func (br *badgeRepository) GetUserCurrentMonthBadge(ctx context.Context, tx *sqlx.Tx, userId int) (Badge, error) {
	executer := br.BaseRepository.initiateQueryExecuter(tx)

	var badge Badge
	err := executer.GetContext(ctx, &badge, getUserCurrentMonthBadgeQuery, userId)
	if err != nil {
		slog.Error("error fetching current month earned badge for user", "error", err)
		return Badge{}, apperrors.ErrBadgeCreationFailed
	}

	return badge, nil

}

func (br *badgeRepository) CreateBadge(ctx context.Context, tx *sqlx.Tx, userId int, badgeType string) (Badge, error) {
	executer := br.BaseRepository.initiateQueryExecuter(tx)

	var createdBadge Badge
	err := executer.GetContext(ctx, &createdBadge, createBadgeQuery, userId, badgeType, time.Now())
	if err != nil {
		slog.Error("error creating badge for user", "error", err)
		return Badge{}, apperrors.ErrBadgeCreationFailed
	}

	return createdBadge, nil
}

func (br *badgeRepository) GetBadgeDetailsOfUser(ctx context.Context, tx *sqlx.Tx, userId int) ([]Badge, error) {
	executer := br.BaseRepository.initiateQueryExecuter(tx)

	var badges []Badge

	err := executer.SelectContext(ctx, &badges, getBadgeDetailsOfUserQuery, userId)
	if err != nil {
		return nil, err
	}

	return badges, nil
}
