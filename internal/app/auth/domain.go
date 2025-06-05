package auth

import (
	"database/sql"
	"time"
)

const (
	LoginWithGithubFailed = "LoginWithGithubFailed"
	AccessTokenCookieName = "AccessToken"
	GitHubOAuthState      = "state"
	GithubOauthScope      = "read:user"
	GetUserGithubUrl      = "https://api.github.com/user"
	GetUserEmailUrl       = "https://api.github.com/user/emails"
)

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
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
}

type GithubUserResponse struct {
	GithubId       int    `json:"id"`
	GithubUsername string `json:"login"`
	AvatarUrl      string `json:"avatar_url"`
	Email          string `json:"email"`
	IsAdmin        bool   `json:"is_admin"`
}
