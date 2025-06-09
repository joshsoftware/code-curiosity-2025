package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type badgeRepository struct {
	BaseRepository
}

type BadgeRepository interface {
	RepositoryTransaction
	GetBadgeDetailsOfUser(ctx context.Context, tx *sqlx.Tx, userId int) ([]Badge, error)
}

func NewBadgeRepository(db *sqlx.DB) BadgeRepository {
	return &badgeRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	getBadgeDetailsOfUserQuery = "SELECT id, badge_type, earned_at FROM badges WHERE user_id = $1"
)

func (br *badgeRepository) GetBadgeDetailsOfUser(ctx context.Context, tx *sqlx.Tx, userId int) ([]Badge, error) {
	executer := br.BaseRepository.initiateQueryExecuter(tx)

	var badges []Badge

	err := executer.SelectContext(ctx, &badges, getBadgeDetailsOfUserQuery, userId)
	if err != nil {
		return nil, err
	}

	return badges, nil
}
