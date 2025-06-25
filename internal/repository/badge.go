package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository/base"
)

type badgeRepository struct {
	base.BaseRepository
}

type BadgeRepository interface {
	base.RepositoryTransaction
	GetBadgeDetailsOfUser(ctx context.Context, tx *sqlx.Tx, userId int) ([]Badge, error)
}

func NewBadgeRepository(db *sqlx.DB) BadgeRepository {
	return &badgeRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}

const (
	getBadgeDetailsOfUserQuery = "SELECT id, badge_type, earned_at FROM badges WHERE user_id = $1 ORDER BY earned_at DESC"
)

func (br *badgeRepository) GetBadgeDetailsOfUser(ctx context.Context, tx *sqlx.Tx, userId int) ([]Badge, error) {
	executer := br.BaseRepository.InitiateQueryExecuter()

	var badges []Badge

	err := executer.SelectContext(ctx, &badges, getBadgeDetailsOfUserQuery, userId)
	if err != nil {
		return nil, err
	}

	return badges, nil
}
