package user

import "database/sql"

type User struct {
	Id                  int           `json:"user_id"`
	GithubId            int           `json:"github_id"`
	GithubUsername      string        `json:"github_username"`
	Email               string        `json:"email"`
	AvatarUrl           string        `json:"avatar_url"`
	CurrentBalance      int           `json:"current_balance"`
	CurrentActiveGoalId sql.NullInt64 `json:"current_active_goal_id"`
	IsBlocked           bool          `json:"is_blocked"`
	IsAdmin             bool          `json:"is_admin"`
	Password            string        `json:"password"`
	IsDeleted           bool          `json:"is_deleted"`
	DeletedAt           sql.NullTime  `json:"deleted_at"`
	CreatedAt           string        `json:"created_at"`
	UpdatedAt           string        `json:"updated_at"`
}

type CreateUserRequestBody struct {
	GithubId       int    `json:"id"`
	GithubUsername string `json:"github_id"`
	AvatarUrl      string `json:"avatar_url"`
	Email          string `json:"email"`
	IsAdmin        bool   `json:"is_admin"`
}

type Email struct {
	Email string `json:"email"`
}
