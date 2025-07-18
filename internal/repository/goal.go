package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type goalRepository struct {
	BaseRepository
}

type GoalRepository interface {
	RepositoryTransaction
	ListGoalLevels(ctx context.Context, tx *sqlx.Tx) ([]Goal, error)
	GetGoalIdByGoalLevel(ctx context.Context, tx *sqlx.Tx, level string) (int, error)
	ListUserGoalLevelTargets(ctx context.Context, tx *sqlx.Tx, userId int) ([]GoalContribution, error)
	CreateCustomGoalLevelTarget(ctx context.Context, tx *sqlx.Tx, customGoalContributionInfo GoalContribution) (GoalContribution, error)
	GetUserActiveGoalLevel(ctx context.Context, tx *sqlx.Tx, userId int) (string, error)
}

func NewGoalRepository(db *sqlx.DB) GoalRepository {
	return &goalRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	listGoalLevelQuery = "SELECT * from goal;"

	getGoalIdByGoalLevelQuery = "SELECT id from goal where level=$1"

	listUserGoalLevelTargetsQuery = `
	SELECT * from goal_contribution 
	where goal_id 
	IN 
	(SELECT current_active_goal_id from users where id=$1)`

	createCustomGoalLevelTargetQuery = `
	INSERT INTO goal_contribution(
	goal_id, 
	contribution_score_id, 
	target_count, 
	is_custom, 
	set_by_user_id
	)
	VALUES 
	($1, $2, $3, $4, $5)
	RETURNING *`

	getUserActiveGoalLevelQuery = `
	SELECT level from goal 
	where id IN 
	(SELECT current_active_goal_id from users where id=$1)`
)

func (gr *goalRepository) ListGoalLevels(ctx context.Context, tx *sqlx.Tx) ([]Goal, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goals []Goal
	err := executer.SelectContext(ctx, &goals, listGoalLevelQuery)
	if err != nil {
		slog.Error("error fetching goal levels", "error", err)
		return nil, apperrors.ErrFetchingGoals
	}

	return goals, nil
}

func (gr *goalRepository) GetGoalIdByGoalLevel(ctx context.Context, tx *sqlx.Tx, level string) (int, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalId int
	err := executer.GetContext(ctx, &goalId, getGoalIdByGoalLevelQuery, level)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("error goal not found", "error", err)
			return 0, apperrors.ErrGoalNotFound
		}

		slog.Error("error occured while getting goal id by goal level", "error", err)
		return 0, apperrors.ErrInternalServer
	}

	return goalId, nil
}

func (gr *goalRepository) ListUserGoalLevelTargets(ctx context.Context, tx *sqlx.Tx, userId int) ([]GoalContribution, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalLevelTargets []GoalContribution
	err := executer.SelectContext(ctx, &goalLevelTargets, listUserGoalLevelTargetsQuery, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("error goal not found", "error", err)
			return nil, apperrors.ErrInternalServer
		}

		slog.Error("error occured while getting goal id by goal level", "error", err)
		return nil, apperrors.ErrInternalServer
	}

	return goalLevelTargets, nil
}

func (gr *goalRepository) CreateCustomGoalLevelTarget(ctx context.Context, tx *sqlx.Tx, customGoalContributionInfo GoalContribution) (GoalContribution, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var customGoalContribution GoalContribution
	err := executer.GetContext(ctx, &customGoalContribution, createCustomGoalLevelTargetQuery,
		customGoalContributionInfo.GoalId,
		customGoalContributionInfo.ContributionScoreId,
		customGoalContributionInfo.TargetCount,
		true,
		customGoalContributionInfo.SetByUserId)
	if err != nil {
		slog.Error("error creating custom goal level targets", "error", err)
		return GoalContribution{}, apperrors.ErrCustomGoalTargetCreationFailed
	}

	return customGoalContribution, nil
}

func (gr *goalRepository) GetUserActiveGoalLevel(ctx context.Context, tx *sqlx.Tx, userId int) (string, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var userActiveGoalLevel string
	err := executer.GetContext(ctx, &userActiveGoalLevel, getUserActiveGoalLevelQuery, userId)
	if err != nil {
		slog.Error("error getting users current active goal level name", "error", err)
		return userActiveGoalLevel, apperrors.ErrInternalServer
	}

	return userActiveGoalLevel, nil
}
