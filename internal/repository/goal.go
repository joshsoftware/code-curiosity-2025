package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

const (
	fetchGoalIdByGoalNameQuery = "SELECT id FROM goal where level ilike '%$1%'"
)

type goalRepository struct {
	BaseRepository
}

type GoalRepository interface {
	RepositoryTransaction
	GetGoalIdByGoalLevel(ctx context.Context, tx *sqlx.Tx, level string) (int, error)
}

func NewGoalRepository(db *sqlx.DB) GoalRepository {
	return &goalRepository{
		BaseRepository: BaseRepository{db},
	}
}

func (gr *goalRepository) GetGoalIdByGoalLevel(ctx context.Context, tx *sqlx.Tx, level string) (int, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalId int
	err := executer.QueryRowContext(ctx, fetchGoalIdByGoalNameQuery, level).Scan(&goalId,)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("[Goal repo] Goal not found", "error", err)
			return 0, apperrors.GoalNotFound
		}

		slog.Error("[Goal repo] Error occured while getting goal id by goal level", "error", err)
		return 0, apperrors.ErrInternalServer
	}

	return goalId, nil
}
