package repository

import (
	"database/sql"
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
