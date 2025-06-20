package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
)

type repositoryRepository struct {
	BaseRepository
}

type RepositoryRepository interface {
	RepositoryTransaction
	GetRepoByGithubId(ctx context.Context, tx *sqlx.Tx, repoGithubId int) (Repository, error)
	GetRepoByRepoId(ctx context.Context, tx *sqlx.Tx, repoId int) (Repository, error)
	CreateRepository(ctx context.Context, tx *sqlx.Tx, repository Repository) (Repository, error)
	GetUserRepoTotalCoins(ctx context.Context, tx *sqlx.Tx, repoId int) (int, error)
	FetchUsersContributedRepos(ctx context.Context, tx *sqlx.Tx) ([]Repository, error)
	FetchUserContributionsInRepo(ctx context.Context, tx *sqlx.Tx, repoGithubId int) ([]Contribution, error)
}

func NewRepositoryRepository(db *sqlx.DB) RepositoryRepository {
	return &repositoryRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	getRepoByGithubIdQuery = `SELECT * from repositories where github_repo_id=$1`

	getrepoByRepoIdQuery = `SELECT * from repositories where id=$1`

	createRepositoryQuery = `
	INSERT INTO repositories (
	github_repo_id, 
	repo_name, 
	description, 
	languages_url,
	repo_url,
	owner_name, 
	update_date,
	contributors_url
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING *`

	getUserRepoTotalCoinsQuery = `SELECT sum(balance_change) from contributions where user_id = $1 and repository_id = $2;`

	fetchUsersContributedReposQuery = `SELECT * from repositories where id in (SELECT repository_id from contributions where user_id=$1);`

	fetchUserContributionsInRepoQuery = `SELECT * from contributions where repository_id in (SELECT id from repositories where github_repo_id=$1) and user_id=$2;`
)

func (rr *repositoryRepository) GetRepoByGithubId(ctx context.Context, tx *sqlx.Tx, repoGithubId int) (Repository, error) {
	executer := rr.BaseRepository.initiateQueryExecuter(tx)

	var repository Repository
	err := executer.QueryRowContext(ctx, getRepoByGithubIdQuery, repoGithubId).Scan(
		&repository.Id,
		&repository.GithubRepoId,
		&repository.RepoName,
		&repository.Description,
		&repository.LanguagesUrl,
		&repository.RepoUrl,
		&repository.OwnerName,
		&repository.UpdateDate,
		&repository.CreatedAt,
		&repository.UpdatedAt,
		&repository.ContributorsUrl,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("repository not found", "error", err)
			return Repository{}, apperrors.ErrRepoNotFound
		}
		slog.Error("error occurred while getting repository by repo github id", "error", err)
		return Repository{}, apperrors.ErrInternalServer
	}

	return repository, nil

}

func (rr *repositoryRepository) GetRepoByRepoId(ctx context.Context, tx *sqlx.Tx, repoId int) (Repository, error) {
	executer := rr.BaseRepository.initiateQueryExecuter(tx)

	var repository Repository
	err := executer.QueryRowContext(ctx, getrepoByRepoIdQuery, repoId).Scan(
		&repository.Id,
		&repository.GithubRepoId,
		&repository.RepoName,
		&repository.Description,
		&repository.LanguagesUrl,
		&repository.RepoUrl,
		&repository.OwnerName,
		&repository.UpdateDate,
		&repository.CreatedAt,
		&repository.UpdatedAt,
		&repository.ContributorsUrl,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Error("repository not found", "error", err)
			return Repository{}, apperrors.ErrRepoNotFound
		}
		slog.Error("error occurred while getting repository by id", "error", err)
		return Repository{}, apperrors.ErrInternalServer
	}

	return repository, nil
}

func (rr *repositoryRepository) CreateRepository(ctx context.Context, tx *sqlx.Tx, repositoryInfo Repository) (Repository, error) {
	executer := rr.BaseRepository.initiateQueryExecuter(tx)

	var repository Repository
	err := executer.QueryRowContext(ctx, createRepositoryQuery,
		repositoryInfo.GithubRepoId,
		repositoryInfo.RepoName,
		repositoryInfo.Description,
		repositoryInfo.LanguagesUrl,
		repositoryInfo.RepoUrl,
		repositoryInfo.OwnerName,
		repositoryInfo.UpdateDate,
		repositoryInfo.ContributorsUrl,
	).Scan(
		&repository.Id,
		&repository.GithubRepoId,
		&repository.RepoName,
		&repository.Description,
		&repository.LanguagesUrl,
		&repository.RepoUrl,
		&repository.OwnerName,
		&repository.UpdateDate,
		&repository.CreatedAt,
		&repository.UpdatedAt,
		&repository.ContributorsUrl,
	)
	if err != nil {
		slog.Error("error occured while creating repository", "error", err)
		return Repository{}, apperrors.ErrInternalServer
	}

	return repository, nil

}

func (r *repositoryRepository) GetUserRepoTotalCoins(ctx context.Context, tx *sqlx.Tx, repoId int) (int, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return 0, apperrors.ErrInternalServer
	}

	executer := r.BaseRepository.initiateQueryExecuter(tx)

	var totalCoins int

	err := executer.QueryRowContext(ctx, getUserRepoTotalCoinsQuery, userId, repoId).Scan(&totalCoins)
	if err != nil {
		slog.Error("error calculating total coins earned by user for the repository")
		return 0, apperrors.ErrCalculatingUserRepoTotalCoins
	}

	return totalCoins, nil
}

func (r *repositoryRepository) FetchUsersContributedRepos(ctx context.Context, tx *sqlx.Tx) ([]Repository, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return nil, apperrors.ErrInternalServer
	}

	executer := r.BaseRepository.initiateQueryExecuter(tx)

	rows, err := executer.QueryContext(ctx, fetchUsersContributedReposQuery, userId)
	if err != nil {
		slog.Error("error fetching users contributed repositories")
		return nil, apperrors.ErrFetchingUsersContributedRepos
	}
	defer rows.Close()

	var usersContributedRepos []Repository
	for rows.Next() {
		var usersContributedRepo Repository
		if err = rows.Scan(
			&usersContributedRepo.Id,
			&usersContributedRepo.GithubRepoId,
			&usersContributedRepo.RepoName,
			&usersContributedRepo.Description,
			&usersContributedRepo.LanguagesUrl,
			&usersContributedRepo.RepoUrl,
			&usersContributedRepo.OwnerName,
			&usersContributedRepo.UpdateDate,
			&usersContributedRepo.CreatedAt,
			&usersContributedRepo.UpdatedAt,
			&usersContributedRepo.ContributorsUrl); err != nil {
			return nil, err
		}

		usersContributedRepos = append(usersContributedRepos, usersContributedRepo)
	}

	return usersContributedRepos, nil
}

func (r *repositoryRepository) FetchUserContributionsInRepo(ctx context.Context, tx *sqlx.Tx, repoGithubId int) ([]Contribution, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return nil, apperrors.ErrInternalServer
	}

	executer := r.BaseRepository.initiateQueryExecuter(tx)

	rows, err := executer.QueryContext(ctx, fetchUserContributionsInRepoQuery, repoGithubId, userId)
	if err != nil {
		slog.Error("error fetching users contribution in repository")
		return nil, apperrors.ErrFetchingUserContributionsInRepo
	}
	defer rows.Close()

	var userContributionsInRepo []Contribution
	for rows.Next() {
		var userContributionInRepo Contribution
		if err = rows.Scan(
			&userContributionInRepo.Id,
			&userContributionInRepo.UserId,
			&userContributionInRepo.RepositoryId,
			&userContributionInRepo.ContributionScoreId,
			&userContributionInRepo.ContributionType,
			&userContributionInRepo.BalanceChange,
			&userContributionInRepo.ContributedAt,
			&userContributionInRepo.CreatedAt,
			&userContributionInRepo.UpdatedAt); err != nil {
			return nil, err
		}

		userContributionsInRepo = append(userContributionsInRepo, userContributionInRepo)
	}

	return userContributionsInRepo, nil
}
