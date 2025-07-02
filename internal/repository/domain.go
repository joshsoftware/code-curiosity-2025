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
	Id                  int       `db:"id"`
	UserId              int       `db:"user_id"`
	RepositoryId        int       `db:"repository_id"`
	ContributionScoreId int       `db:"contribution_score_id"`
	ContributionType    string    `db:"contribution_type"`
	BalanceChange       int       `db:"balance_change"`
	ContributedAt       time.Time `db:"contributed_at"`
	GithubEventId       string    `db:"github_event_id"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type Repository struct {
	Id              int       `db:"id"`
	GithubRepoId    int       `db:"github_repo_id"`
	RepoName        string    `db:"repo_name"`
	Description     string    `db:"description"`
	LanguagesUrl    string    `db:"languages_url"`
	RepoUrl         string    `db:"repo_url"`
	OwnerName       string    `db:"owner_name"`
	UpdateDate      time.Time `db:"update_date"`
	ContributorsUrl string    `db:"contributors_url"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type ContributionScore struct {
	Id               int       `db:"id"`
	AdminId          int       `db:"admin_id"`
	ContributionType string    `db:"contribution_type"`
	Score            int       `db:"score"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}
