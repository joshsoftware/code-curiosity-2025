package repository

import (
	"database/sql"
	"time"
)

type User struct {
	Id                  int           `db:"id"`
	GithubId            int           `db:"github_id"`
	GithubUsername      string        `db:"github_username"`
	Email               string        `db:"email"`
	AvatarUrl           string        `db:"avatar_url"`
	CurrentBalance      int           `db:"current_balance"`
	CurrentActiveGoalId sql.NullInt64 `db:"current_active_goal_id"`
	IsBlocked           bool          `db:"is_blocked"`
	IsAdmin             bool          `db:"is_admin"`
	Password            string        `db:"password"`
	IsDeleted           bool          `db:"is_deleted"`
	DeletedAt           sql.NullTime  `db:"deleted_at"`
	CreatedAt           time.Time     `db:"created_at"`
	UpdatedAt           time.Time     `db:"updated_at"`
}

type CreateUserRequestBody struct {
	GithubId       int    `db:"github_id"`
	GithubUsername string `db:"github_username"`
	AvatarUrl      string `db:"avatar_url"`
	Email          string `db:"email"`
	IsAdmin        bool   `db:"is_admin"`
}

type Contribution struct {
	Id                  int
	UserId              int
	RepositoryId        int
	ContributionScoreId int
	ContributionType    string
	BalanceChange       int
	ContributedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Repository struct {
	Id              int
	GithubRepoId    int
	RepoName        string
	Description     string
	LanguagesUrl    string
	RepoUrl         string
	OwnerName       string
	UpdateDate      time.Time
	ContributorsUrl string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ContributionScore struct {
	Id               int
	AdminId          int
	ContributionType string
	Score            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
