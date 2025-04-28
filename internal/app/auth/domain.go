package auth

const (
	LoginWithGithubFailed = "LoginWithGithubFailed"
	AccessTokenCookieName = "AccessToken"
	GetUserGithubUrl      = "https://api.github.com/user"
	GetUserEmailUrl       = "https://api.github.com/user/emails"
)

type User struct {
	UserId         int    `json:"user_id"`
	GithubId       int    `json:"id"`
	GithubUsername string `json:"login"`
	AvatarUrl      string `json:"avatar_url"`
	Email          string `json:"email"`
	IsAdmin        bool   `json:"is_admin"`
}
