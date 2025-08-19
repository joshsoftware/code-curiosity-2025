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
	Id                  int           `json:"userId"`
	GithubId            int           `json:"githubId"`
	GithubUsername      string        `json:"githubUsername"`
	Email               string        `json:"email"`
	AvatarUrl           string        `json:"avatarUrl"`
	CurrentBalance      int           `json:"currentBalance"`
	CurrentActiveGoalId sql.NullInt64 `json:"currentActiveGoalId"`
	IsBlocked           bool          `json:"isBlocked"`
	IsAdmin             bool          `json:"isAdmin"`
	Password            string        `json:"password"`
	IsDeleted           bool          `json:"isDeleted"`
	DeletedAt           sql.NullTime  `json:"deletedAt"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
}

type GithubUserResponse struct {
	GithubId       int    `json:"id"`
	GithubUsername string `json:"login"`
	AvatarUrl      string `json:"avatar_url"`
	Email          string `json:"email"`
	IsAdmin        bool   `json:"is_admin"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	User
	JwtToken string `json:"jwtToken"`
}
