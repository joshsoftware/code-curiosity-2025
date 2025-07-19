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

type userRepository struct {
	BaseRepository
}

type UserRepository interface {
	RepositoryTransaction
	GetUserById(ctx context.Context, tx *sqlx.Tx, userId int) (User, error)
	GetUserByGithubId(ctx context.Context, tx *sqlx.Tx, githubId int) (User, error)
	CreateUser(ctx context.Context, tx *sqlx.Tx, userInfo CreateUserRequestBody) (User, error)
	UpdateUserEmail(ctx context.Context, tx *sqlx.Tx, userId int, email string) error
	MarkUserAsDeleted(ctx context.Context, tx *sqlx.Tx, userID int, deletedAt time.Time) error
	RecoverAccountInGracePeriod(ctx context.Context, tx *sqlx.Tx, userID int) error
	HardDeleteUsers(ctx context.Context, tx *sqlx.Tx) error
	GetAllUsersGithubId(ctx context.Context, tx *sqlx.Tx) ([]int, error)
	UpdateUserCurrentBalance(ctx context.Context, tx *sqlx.Tx, user User) error
	GetAllUsersRank(ctx context.Context, tx *sqlx.Tx) ([]LeaderboardUser, error)
	GetCurrentUserRank(ctx context.Context, tx *sqlx.Tx, userId int) (LeaderboardUser, error)
	UpdateCurrentActiveGoalId(ctx context.Context, tx *sqlx.Tx, userId int, goalId int) (int, error)
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	getUserByIdQuery = "SELECT * from users where id=$1"

	getUserByGithubIdQuery = "SELECT * from users where github_id=$1"

	createUserQuery = `
	INSERT INTO users ( 
	github_id, 
	github_username, 
	email, 
	avatar_url
	) 
	VALUES ($1, $2, $3, $4) 
	RETURNING *`

	updateEmailQuery = "UPDATE users SET email=$1, updated_at=$2 where id=$3"

	markUserAsDeletedQuery = "UPDATE users SET is_deleted = TRUE, deleted_at=$1 where id = $2"

	recoverAccountInGracePeriodQuery = "UPDATE users SET is_deleted = false, deleted_at = NULL where id = $1"

	hardDeleteUsersQuery = "DELETE FROM users WHERE is_deleted = TRUE AND deleted_at <= $1"

	getAllUsersGithubIdQuery = "SELECT github_id from users"

	updateUserCurrentBalanceQuery = "UPDATE users SET current_balance=$1, updated_at=$2 where id=$3"

	getAllUsersRankQuery = `
	SELECT 
	id,
	github_username,
	avatar_url,
	current_balance,
	RANK() over (ORDER BY current_balance DESC) AS rank
	FROM users 
	ORDER BY current_balance DESC`

	getCurrentUserRankQuery = `
	SELECT *
	FROM 
	(
  	SELECT 
	id, 
	github_username, 
	avatar_url,
	current_balance,
    RANK() OVER (ORDER BY current_balance DESC) AS rank
  	FROM users
	) 
	ranked_users
	WHERE id = $1;`

	updateCurrentActiveGoalIdQuery = "UPDATE users SET current_active_goal_id=$1 where id=$2"
)

func (ur *userRepository) GetUserById(ctx context.Context, tx *sqlx.Tx, userId int) (User, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var user User
	err := executer.GetContext(ctx, &user, getUserByIdQuery, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("user not found", "error", err)
			return User{}, apperrors.ErrUserNotFound
		}
		slog.Error("error occurred while getting user by id", "error", err)
		return User{}, apperrors.ErrInternalServer
	}

	return user, nil
}

func (ur *userRepository) GetUserByGithubId(ctx context.Context, tx *sqlx.Tx, githubId int) (User, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var user User
	err := executer.GetContext(ctx, &user, getUserByGithubIdQuery, githubId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("user not found", "error", err)
			return User{}, apperrors.ErrUserNotFound
		}
		slog.Error("error occurred while getting user by github id", "error", err)
		return User{}, apperrors.ErrInternalServer
	}

	return user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, tx *sqlx.Tx, userInfo CreateUserRequestBody) (User, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var user User
	err := executer.GetContext(ctx, &user, createUserQuery,
		userInfo.GithubId,
		userInfo.GithubUsername,
		userInfo.Email,
		userInfo.AvatarUrl)

	if err != nil {
		slog.Error("error occurred while creating user", "error", err)
		return User{}, apperrors.ErrUserCreationFailed
	}

	return user, nil

}

func (ur *userRepository) UpdateUserEmail(ctx context.Context, tx *sqlx.Tx, userId int, email string) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, updateEmailQuery, email, time.Now(), userId)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		return apperrors.ErrInternalServer
	}

	return nil
}

func (ur *userRepository) MarkUserAsDeleted(ctx context.Context, tx *sqlx.Tx, userID int, deletedAt time.Time) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, markUserAsDeletedQuery, deletedAt, userID)
	if err != nil {
		slog.Error("unable to mark user as deleted", "error", err)
		return apperrors.ErrInternalServer
	}

	return nil
}

func (ur *userRepository) RecoverAccountInGracePeriod(ctx context.Context, tx *sqlx.Tx, userID int) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, recoverAccountInGracePeriodQuery, userID)
	if err != nil {
		slog.Error("unable to reverse the soft delete ", "error", err)
		return apperrors.ErrInternalServer
	}

	return nil
}

func (ur *userRepository) HardDeleteUsers(ctx context.Context, tx *sqlx.Tx) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	threshold := time.Now().Add(-90 * 1 * time.Second)

	_, err := executer.ExecContext(ctx, hardDeleteUsersQuery, threshold)
	if err != nil {
		slog.Error("error deleting users that are soft deleted for more than three months", "error", err)
		return apperrors.ErrInternalServer
	}

	return err
}

func (ur *userRepository) GetAllUsersGithubId(ctx context.Context, tx *sqlx.Tx) ([]int, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var githubIds []int
	err := executer.SelectContext(ctx, &githubIds, getAllUsersGithubIdQuery)
	if err != nil {
		slog.Error("failed to get github usernames", "error", err)
		return nil, apperrors.ErrInternalServer
	}

	return githubIds, nil
}

func (ur *userRepository) UpdateUserCurrentBalance(ctx context.Context, tx *sqlx.Tx, user User) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, updateUserCurrentBalanceQuery, user.CurrentBalance, time.Now(), user.Id)
	if err != nil {
		slog.Error("failed to update user balance change", "error", err)
		return apperrors.ErrInternalServer
	}

	return nil
}

func (ur *userRepository) GetAllUsersRank(ctx context.Context, tx *sqlx.Tx) ([]LeaderboardUser, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var leaderboard []LeaderboardUser
	err := executer.SelectContext(ctx, &leaderboard, getAllUsersRankQuery)
	if err != nil {
		slog.Error("failed to get users rank", "error", err)
		return nil, apperrors.ErrInternalServer
	}

	return leaderboard, nil
}

func (ur *userRepository) GetCurrentUserRank(ctx context.Context, tx *sqlx.Tx, userId int) (LeaderboardUser, error) {

	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var currentUserRank LeaderboardUser
	err := executer.GetContext(ctx, &currentUserRank, getCurrentUserRankQuery, userId)
	if err != nil {
		slog.Error("failed to get user rank", "error", err)
		return LeaderboardUser{}, apperrors.ErrInternalServer
	}

	return currentUserRank, nil
}

func (ur *userRepository) UpdateCurrentActiveGoalId(ctx context.Context, tx *sqlx.Tx, userId int, goalId int) (int, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, updateCurrentActiveGoalIdQuery, goalId, userId)
	if err != nil {
		slog.Error("failed to update current active goal id", "error", err)
		return 0, apperrors.ErrInternalServer
	}

	return goalId, nil
}
