package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type goalRepository struct {
	BaseRepository
}

type GoalRepository interface {
	RepositoryTransaction
	ListGoalLevels(ctx context.Context, tx *sqlx.Tx) ([]GoalLevel, error)
	GetGoalLevelByLevel(ctx context.Context, tx *sqlx.Tx, level string) (GoalLevel, error)
	CreateUserGoalInProgress(ctx context.Context, tx *sqlx.Tx, userGoal UserGoal) (UserGoal, error)
	FetchGoalLevelTargetByGoalLevel(ctx context.Context, tx *sqlx.Tx, goalLevel GoalLevel) ([]GoalLevelTarget, error)
	CreateUserGoalTarget(ctx context.Context, tx *sqlx.Tx, userGoalTarget UserGoalTarget) (UserGoalTarget, error)
	CreateUserGoalProgress(ctx context.Context, tx *sqlx.Tx, userGoalProgress UserGoalProgress) (UserGoalProgress, error)
	GetUserCurrentGoal(ctx context.Context, tx *sqlx.Tx, userId int) (UserGoal, error)
	UpdateUserGoalStatus(ctx context.Context, tx *sqlx.Tx, userGoal UserGoal) (UserGoal, error)
	ListUserGoalTargetsByUserGoalId(ctx context.Context, tx *sqlx.Tx, userGoalId int) ([]UserGoalTarget, error)
	GetContributionProgressCount(ctx context.Context, tx *sqlx.Tx, userGoalTargetId int) (int, error)
	GetGoalLevelById(ctx context.Context, tx *sqlx.Tx, goalLevelId int) (GoalLevel, error)
	FetchInProgressUserGoalsOfPreviousMonth(ctx context.Context, tx *sqlx.Tx) ([]UserGoal, error)
	CalculateUserIncompleteGoalsUntilDay(ctx context.Context, tx *sqlx.Tx, userID int) (int, error)
	CreateUserGoalSummary(ctx context.Context, tx *sqlx.Tx, userGoalSummary GoalSummary) (GoalSummary, error)
	FetchUserGoalSummary(ctx context.Context, tx *sqlx.Tx, userId int) ([]GoalSummary, error)
	GetUserGoalSummaryBySnapshotDate(ctx context.Context, tx *sqlx.Tx, snapshotDate time.Time, userId int) (*GoalSummary, error)
	UpdateUserGoalSummary(ctx context.Context, tx *sqlx.Tx, goalSummaryId int, userGoalSummary GoalSummary) (GoalSummary, error)
}

func NewGoalRepository(db *sqlx.DB) GoalRepository {
	return &goalRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	listGoalLevelQuery = "SELECT * from goal_level;"

	getGoalLevelByLevelQuery = "SELECT * from goal_level where level=$1"

	createUserGoalInProgressQuery = `
	INSERT INTO user_goal(
	user_id,
	goal_level_id,
	status,
	month_started_at
	)
	VALUES
	($1, $2, $3, $4)
	RETURNING *`

	fetchGoalLevelTargetByGoalLevelQuery = "SELECT * FROM goal_level_target WHERE goal_level_id=$1"

	createUserGoalTargetQuery = `
	INSERT INTO user_goal_target(
	user_goal_id,
	contribution_score_id,
	target
	)
	VALUES
	($1, $2, $3)
	RETURNING *`

	createUserGoalProgressQuery = `
	INSERT INTO user_goal_progress(
	user_goal_target_id,
	contribution_id
	)
	VALUES
	($1, $2)
	RETURNING *`

	getUserCurrentGoalQuery = `
	SELECT * FROM user_goal
	WHERE date_trunc('month', month_started_at) = date_trunc('month', NOW())
	AND status != 'incomplete'
	AND user_id = $1
	ORDER BY created_at DESC
	LIMIT 1`

	updateUserGoalStatusQuery = "UPDATE user_goal SET status=$1, updated_at=$2 WHERE id=$3 RETURNING *"

	listUserGoalTargetsByUserGoalIdQuery = "SELECT * from user_goal_target WHERE user_goal_id=$1"

	getContributionProgressCountQuery = "SELECT COUNT(*)  FROM user_goal_progress WHERE user_goal_target_id=$1"

	getGoalLevelByIdQuery = "SELECT * FROM goal_level where id=$1"

	fetchInProgressUserGoalsOfPreviousMonthQuery = `
	SELECT * FROM user_goal
	WHERE date_trunc('month', month_started_at) = date_trunc('month', NOW() - interval '1 month')
	WHERE status='inProgress'`

	calculateUserIncompleteGoalsUntilDayQuery = `
	SELECT COUNT(*) FROM user_goal
	WHERE date_trunc('month', month_started_at) = date_trunc('month', NOW())
	AND status = 'incomplete'
	AND user_id = $1`

	createUserGoalSummaryQuery = `
	INSERT INTO goal_summary(
	user_id,
	snapshot_date,
	incomplete_goals_count,
	target_set,
	target_completed
	)
	VALUES
	($1, $2, $3, $4, $5)
	RETURNING *`

	fetchUserGoalSummaryQuery = "SELECT * FROM goal_summary WHERE user_id=$1"

	getUserGoalSummaryBySnapshotDateQuery = `
	SELECT * FROM goal_summary
	WHERE user_id = $2
  	AND snapshot_date::date = $1::date
	LIMIT 1`

	updateUserGoalSummaryQuery = "UPDATE goal_summary SET snapshot_date=$2, incomplete_goals_count=$3, target_set=$4, target_completed=$5 where id=$1 "
)

func (gr *goalRepository) ListGoalLevels(ctx context.Context, tx *sqlx.Tx) ([]GoalLevel, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalLevels []GoalLevel
	err := executer.SelectContext(ctx, &goalLevels, listGoalLevelQuery)
	if err != nil {
		slog.Error("error fetching goal levels", "error", err)
		return nil, apperrors.ErrFetchingGoals
	}

	return goalLevels, nil
}

func (gr *goalRepository) GetGoalLevelByLevel(ctx context.Context, tx *sqlx.Tx, level string) (GoalLevel, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalLevel GoalLevel
	err := executer.GetContext(ctx, &goalLevel, getGoalLevelByLevelQuery, level)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("error goal level not found", "error", err)
			return GoalLevel{}, apperrors.ErrGoalLevelNotFound
		}

		slog.Error("error occured while getting goal level by level", "error", err)
		return GoalLevel{}, apperrors.ErrFailedToGetGoalLevel
	}

	return goalLevel, nil
}

func (gr *goalRepository) CreateUserGoalInProgress(ctx context.Context, tx *sqlx.Tx, userGoal UserGoal) (UserGoal, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var createdUserGoal UserGoal
	err := executer.GetContext(ctx, &createdUserGoal, createUserGoalInProgressQuery, userGoal.UserId, userGoal.GoalLevelId, userGoal.Status, userGoal.MonthStartedAt)
	if err != nil {
		slog.Error("failed to create user goal in progress", "error", err)
		return UserGoal{}, apperrors.ErrUserGoalCreationFailed
	}

	return createdUserGoal, nil
}

func (gr *goalRepository) FetchGoalLevelTargetByGoalLevel(ctx context.Context, tx *sqlx.Tx, goalLevel GoalLevel) ([]GoalLevelTarget, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalLevelTargets []GoalLevelTarget
	err := executer.SelectContext(ctx, &goalLevelTargets, fetchGoalLevelTargetByGoalLevelQuery, goalLevel.Id)
	if err != nil {
		slog.Error("error fetching goal level target by goal level", "error", err)
		return nil, apperrors.ErrFetchingGoalLevelTargets
	}

	return goalLevelTargets, nil
}

func (gr *goalRepository) CreateUserGoalTarget(ctx context.Context, tx *sqlx.Tx, userGoalTarget UserGoalTarget) (UserGoalTarget, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var createdUserGoalTarget UserGoalTarget
	err := executer.GetContext(ctx, &createdUserGoalTarget, createUserGoalTargetQuery, userGoalTarget.UserGoalId, userGoalTarget.ContributionScoreId, userGoalTarget.Target)
	if err != nil {
		slog.Error("error creating user goal target", "error", err)
		return UserGoalTarget{}, apperrors.ErrUserGoalTargetCreationFailed
	}

	return createdUserGoalTarget, nil
}

func (gr *goalRepository) CreateUserGoalProgress(ctx context.Context, tx *sqlx.Tx, userGoalProgress UserGoalProgress) (UserGoalProgress, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var createdUserGoalProgress UserGoalProgress
	err := executer.GetContext(ctx, &createdUserGoalProgress, createUserGoalProgressQuery, userGoalProgress.UserGoalTargetId, userGoalProgress.ContributionId)
	if err != nil {
		slog.Error("error creating user goal target", "error", err)
		return UserGoalProgress{}, apperrors.ErrUserGoalProgressCreationFailed
	}

	return createdUserGoalProgress, nil
}

func (gr *goalRepository) GetUserCurrentGoal(ctx context.Context, tx *sqlx.Tx, userId int) (UserGoal, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var userGoal UserGoal
	err := executer.GetContext(ctx, &userGoal, getUserCurrentGoalQuery, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("no user goal found for current month", "userId", userId)
			return UserGoal{}, apperrors.ErrUserGoalNotFound
		}

		slog.Error("error occurred while getting latest user goal for current month", "error", err)
		return UserGoal{}, apperrors.ErrFailedToGetUserGoal
	}

	return userGoal, nil
}

func (gr *goalRepository) UpdateUserGoalStatus(ctx context.Context, tx *sqlx.Tx, userGoal UserGoal) (UserGoal, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var updatedGoal UserGoal
	err := executer.GetContext(ctx, &updatedGoal, updateUserGoalStatusQuery,
		userGoal.Status,
		time.Now().UTC(),
		userGoal.Id,
	)
	if err != nil {
		slog.Error("failed to update user goal id", "error", err)
		return UserGoal{}, apperrors.ErrInternalServer
	}

	return updatedGoal, nil
}

func (gr *goalRepository) ListUserGoalTargetsByUserGoalId(ctx context.Context, tx *sqlx.Tx, userGoalId int) ([]UserGoalTarget, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var userGoalTargets []UserGoalTarget
	err := executer.SelectContext(ctx, &userGoalTargets, listUserGoalTargetsByUserGoalIdQuery, userGoalId)
	if err != nil {
		slog.Error("error fetching user goal targets by user goal id", "error", err)
		return nil, apperrors.ErrFetchingGoalLevelTargets
	}

	return userGoalTargets, nil
}

func (gr *goalRepository) GetContributionProgressCount(ctx context.Context, tx *sqlx.Tx, userGoalTargetId int) (int, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var contributionProgressCount int
	err := executer.GetContext(ctx, &contributionProgressCount, getContributionProgressCountQuery, userGoalTargetId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("no contribution for the target")
			return 0, nil
		}

		slog.Error("error counting progress of given contribution target", "error", err)
		return 0, apperrors.ErrInternalServer
	}

	return contributionProgressCount, nil
}

func (gr *goalRepository) GetGoalLevelById(ctx context.Context, tx *sqlx.Tx, goalLevelId int) (GoalLevel, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var goalLevel GoalLevel
	err := executer.GetContext(ctx, &goalLevel, getGoalLevelByIdQuery, goalLevelId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("error goal level not found", "error", err)
			return GoalLevel{}, apperrors.ErrGoalLevelNotFound
		}

		slog.Error("error occured while getting goal level by id", "error", err)
		return GoalLevel{}, apperrors.ErrFailedToGetGoalLevel
	}

	return goalLevel, nil
}

func (gr *goalRepository) FetchInProgressUserGoalsOfPreviousMonth(ctx context.Context, tx *sqlx.Tx) ([]UserGoal, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var userGoals []UserGoal
	err := executer.SelectContext(ctx, &userGoals, fetchInProgressUserGoalsOfPreviousMonthQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Info("no in progress goals")
			return nil, nil
		}
		slog.Error("error occurred while fetching in progress user goals for previous month", "error", err)
		return nil, apperrors.ErrFailedToGetUserGoal
	}

	return userGoals, nil
}

func (gr *goalRepository) CalculateUserIncompleteGoalsUntilDay(ctx context.Context, tx *sqlx.Tx, userID int) (int, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var userIncompleteGoalCount int
	err := executer.GetContext(ctx, &userIncompleteGoalCount, calculateUserIncompleteGoalsUntilDayQuery, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Info("no incomplete goals for month")
			return 0, nil
		}
		slog.Error("error getting incomplete goals until day", "error", err)
		return 0, err
	}

	return userIncompleteGoalCount, nil
}

func (gr *goalRepository) CreateUserGoalSummary(ctx context.Context, tx *sqlx.Tx, userGoalSummary GoalSummary) (GoalSummary, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var createdUserGoalSummary GoalSummary
	err := executer.GetContext(ctx, &createdUserGoalSummary, createUserGoalSummaryQuery, userGoalSummary.UserId, userGoalSummary.SnapshotDate, userGoalSummary.IncompleteGoalsCount, userGoalSummary.TargetSet, userGoalSummary.TargetCompleted)
	if err != nil {
		slog.Error("error creating user goal summary", "error", err)
		return GoalSummary{}, err
	}

	return createdUserGoalSummary, nil
}

func (gr *goalRepository) FetchUserGoalSummary(ctx context.Context, tx *sqlx.Tx, userId int) ([]GoalSummary, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var usersGoalSummary []GoalSummary
	err := executer.SelectContext(ctx, &usersGoalSummary, fetchUserGoalSummaryQuery, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		slog.Error("error fetching users goal summary", "error", err)
		return nil, err
	}

	return usersGoalSummary, nil
}

func (gr *goalRepository) GetUserGoalSummaryBySnapshotDate(ctx context.Context, tx *sqlx.Tx, snapshotDate time.Time, userId int) (*GoalSummary, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var userGoalSummary GoalSummary
	err := executer.GetContext(ctx, &userGoalSummary, getUserGoalSummaryBySnapshotDateQuery, snapshotDate, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		slog.Error("error getting user goal summary", "error", err)
		return nil, err
	}

	return &userGoalSummary, nil
}

func (gr *goalRepository) UpdateUserGoalSummary(ctx context.Context, tx *sqlx.Tx, goalSummaryId int, userGoalSummary GoalSummary) (GoalSummary, error) {
	executer := gr.BaseRepository.initiateQueryExecuter(tx)

	var updatedUserGoalSummary GoalSummary
	err := executer.GetContext(ctx, &updatedUserGoalSummary, updateUserGoalSummaryQuery,
		goalSummaryId,
		userGoalSummary.SnapshotDate,
		userGoalSummary.IncompleteGoalsCount,
		userGoalSummary.TargetSet,
		userGoalSummary.TargetCompleted,
	)
	if err != nil {
		slog.Error("failed to update user goal summary", "error", err)
		return GoalSummary{}, apperrors.ErrInternalServer
	}

	return GoalSummary{}, nil
}
