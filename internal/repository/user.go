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
	ListContributionScores(ctx context.Context, tx *sqlx.Tx) ([]ContributionScore, error)
	UpdateContributionScore(ctx context.Context, tx *sqlx.Tx, contributionScoreId int, score int) error
	ListLeaderboard(ctx context.Context, tx *sqlx.Tx) ([]LeaderboardEntry, error)
	UpdateUserBlocked(ctx context.Context, tx *sqlx.Tx, userId int, isBlocked bool) error
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
	listContributionScoresQuery = "SELECT id, admin_id, contribution_type, score, created_at, updated_at FROM contribution_score ORDER BY contribution_type"
	updateContributionScoreQuery = "UPDATE contribution_score SET score=$1, updated_at=$2 WHERE id=$3"
	listLeaderboardQuery = "SELECT id, user_id, github_id, avatar_url, current_balance, rank, refreshed_at, created_at, updated_at FROM leaderboard_hourly WHERE refreshed_at = (SELECT MAX(refreshed_at) FROM leaderboard_hourly) ORDER BY rank"
	updateUserBlockedQuery = "UPDATE users SET is_blocked=$1, updated_at=$2 WHERE id=$3 AND is_admin=FALSE"
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

func (ur *userRepository) UpdateUserEmail(ctx context.Context, tx *sqlx.Tx, userId int, email string) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	_, err := executer.ExecContext(ctx, updateEmailQuery, email, time.Now(), userId)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		return apperrors.ErrInternalServer
	}

	return nil
}

func (ur *userRepository) ListContributionScores(ctx context.Context, tx *sqlx.Tx) ([]ContributionScore, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	rows, err := executer.QueryContext(ctx, listContributionScoresQuery)
	if err != nil {
		slog.Error("failed to list contribution scores", "error", err)
		return nil, apperrors.ErrInternalServer
	}
	defer rows.Close()

	contributionScores := []ContributionScore{}
	for rows.Next() {
		var contributionScore ContributionScore
		err = rows.Scan(
			&contributionScore.Id,
			&contributionScore.AdminId,
			&contributionScore.ContributionType,
			&contributionScore.Score,
			&contributionScore.CreatedAt,
			&contributionScore.UpdatedAt,
		)
		if err != nil {
			slog.Error("failed to scan contribution score", "error", err)
			return nil, apperrors.ErrInternalServer
		}
		contributionScores = append(contributionScores, contributionScore)
	}
	if err = rows.Err(); err != nil {
		slog.Error("failed to read contribution scores", "error", err)
		return nil, apperrors.ErrInternalServer
	}

	return contributionScores, nil
}

func (ur *userRepository) UpdateContributionScore(ctx context.Context, tx *sqlx.Tx, contributionScoreId int, score int) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	result, err := executer.ExecContext(ctx, updateContributionScoreQuery, score, time.Now(), contributionScoreId)
	if err != nil {
		slog.Error("failed to update contribution score", "error", err)
		return apperrors.ErrInternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("failed to check updated contribution score", "error", err)
		return apperrors.ErrInternalServer
	}
	if rowsAffected == 0 {
		return apperrors.ErrInvalidQueryParams
	}

	return nil
}

func (ur *userRepository) ListLeaderboard(ctx context.Context, tx *sqlx.Tx) ([]LeaderboardEntry, error) {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	rows, err := executer.QueryContext(ctx, listLeaderboardQuery)
	if err != nil {
		slog.Error("failed to list leaderboard", "error", err)
		return nil, apperrors.ErrInternalServer
	}
	defer rows.Close()

	leaderboard := []LeaderboardEntry{}
	for rows.Next() {
		var leaderboardEntry LeaderboardEntry
		err = rows.Scan(
			&leaderboardEntry.Id,
			&leaderboardEntry.UserId,
			&leaderboardEntry.GithubId,
			&leaderboardEntry.AvatarUrl,
			&leaderboardEntry.CurrentBalance,
			&leaderboardEntry.Rank,
			&leaderboardEntry.RefreshedAt,
			&leaderboardEntry.CreatedAt,
			&leaderboardEntry.UpdatedAt,
		)
		if err != nil {
			slog.Error("failed to scan leaderboard", "error", err)
			return nil, apperrors.ErrInternalServer
		}
		leaderboard = append(leaderboard, leaderboardEntry)
	}
	if err = rows.Err(); err != nil {
		slog.Error("failed to read leaderboard", "error", err)
		return nil, apperrors.ErrInternalServer
	}

	return leaderboard, nil
}

func (ur *userRepository) UpdateUserBlocked(ctx context.Context, tx *sqlx.Tx, userId int, isBlocked bool) error {
	executer := ur.BaseRepository.initiateQueryExecuter(tx)

	result, err := executer.ExecContext(ctx, updateUserBlockedQuery, isBlocked, time.Now(), userId)
	if err != nil {
		slog.Error("failed to update user blocked status", "error", err)
		return apperrors.ErrInternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("failed to check updated user blocked status", "error", err)
		return apperrors.ErrInternalServer
	}
	if rowsAffected == 0 {
		return apperrors.ErrUserNotFound
	}

	return nil
}
