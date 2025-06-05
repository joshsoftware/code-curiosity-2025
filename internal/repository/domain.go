package repository

import (
	"database/sql"
	"time"
)

type User struct {
	Id                  int
	GithubId            int
	GithubUsername      string
	Email               string
	AvatarUrl           string
	CurrentBalance      int
	CurrentActiveGoalId sql.NullInt64
	IsBlocked           bool
	IsAdmin             bool
	Password            string
	IsDeleted           bool
	DeletedAt           sql.NullInt64
	CreatedAt           int64
	UpdatedAt           int64
}

type CreateUserRequestBody struct {
	GithubId       int
	GithubUsername string
	AvatarUrl      string
	Email          string
	IsAdmin        bool
}

type Contribution struct {
	Id                  int
	UserId              int
	RepositoryId        int
	ContributionScoreId int
	ContributionType    string
	BalanceChange       int
	ContributedAt       time.Time
	CreatedAt           int64
	UpdatedAt           int64
}

type Repository struct {
	Id           int
	GithubRepoId int
	RepoName     string
	Description  string
	LanguagesUrl string
	RepoUrl      string
	OwnerName    string
	UpdateDate   time.Time
	CreatedAt    int64
	UpdatedAt    int64
}
