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

type Badge struct {
	BadgeId   int       `db:"id"`
	BadgeType string    `db:"badge_type"`
	EarnedAt  time.Time `db:"earned_at"`
}
