package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type repositoryRepository struct {
	BaseRepository
}

type RepositoryRepository interface {
	RepositoryTransaction
	GetRepoByGithubId(ctx context.Context, tx *sqlx.Tx, repoGithubId int) (Repository, error)
	CreateRepository(ctx context.Context, tx *sqlx.Tx, repository Repository) (Repository, error)
}

func NewRepositoryRepository(db *sqlx.DB) RepositoryRepository {
	return &repositoryRepository{
		BaseRepository: BaseRepository{db},
	}
}

const (
	getRepoByGithubIdQuery = `SELECT * from repositories where github_repo_id=$1`

	createRepositoryQuery = `
	INSERT INTO repositories (
	github_repo_id, 
	repo_name, 
	description, 
	languages_url,
	repo_url,
	owner_name, 
	update_date
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING *`
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
	)
	if err != nil {
		slog.Error("error occured while creating repository", "error", err)
		return Repository{}, apperrors.ErrInternalServer
	}

	return repository, nil

}
