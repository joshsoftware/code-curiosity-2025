package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

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

	updateEmailQuery = "UPDATE users SET email=$1 where id=$2"
)

func (ur *userRepository) GetUserById(ctx context.Context, tx *sqlx.Tx, userId int) (User, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	var user User
	err := executer.QueryRowContext(ctx, getUserByIdQuery, userId).Scan(
		&user.Id,
		&user.GithubId,
		&user.GithubUsername,
		&user.AvatarUrl,
		&user.Email,
		&user.CurrentActiveGoalId,
		&user.CurrentBalance,
		&user.IsBlocked,
		&user.IsAdmin,
		&user.Password,
		&user.IsDeleted,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
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
	err := executer.QueryRowContext(ctx, getUserByGithubIdQuery, githubId).Scan(
		&user.Id,
		&user.GithubId,
		&user.GithubUsername,
		&user.AvatarUrl,
		&user.Email,
		&user.CurrentActiveGoalId,
		&user.CurrentBalance,
		&user.IsBlocked,
		&user.IsAdmin,
		&user.Password,
		&user.IsDeleted,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
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
	err := executer.QueryRowContext(ctx, createUserQuery,
		userInfo.GithubId,
		userInfo.GithubUsername,
		userInfo.Email,
		userInfo.AvatarUrl,
	).Scan(
		&user.Id,
		&user.GithubId,
		&user.GithubUsername,
		&user.AvatarUrl,
		&user.Email,
		&user.CurrentActiveGoalId,
		&user.CurrentBalance,
		&user.IsBlocked,
		&user.IsAdmin,
		&user.Password,
		&user.IsDeleted,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		slog.Error("error occurred while creating user", "error", err)
		return User{}, apperrors.ErrUserCreationFailed
	}

	return user, nil

}

func (ur *userRepository) UpdateUserEmail(ctx context.Context, tx *sqlx.Tx, Id int, email string) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, updateEmailQuery, email, Id)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		return apperrors.ErrInternalServer
	}

	return nil
}
