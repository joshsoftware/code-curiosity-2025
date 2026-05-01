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
	DeletedAt           sql.NullTime
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type CreateUserRequestBody struct {
	GithubId       int
	GithubUsername string
	AvatarUrl      string
	Email          string
	IsAdmin        bool
}

type ContributionScore struct {
	Id               int
	AdminId          int
	ContributionType string
	Score            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type LeaderboardEntry struct {
	Id             int
	UserId         int
	GithubId       int
	AvatarUrl      string
	CurrentBalance int
	Rank           int
	RefreshedAt    time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
