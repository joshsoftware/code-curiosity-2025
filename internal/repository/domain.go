package repository

type User struct {
	Id             int
	GithubId       int
	GithubUsername string
	Email          string
	AvatarUrl      string
	CurrentBalance int
	IsBlocked      bool
	IsAdmin        bool
	Password       string
	CreatedAt      string
	UpdatedAt      string
}

type CreateUserRequestBody struct {
	GithubId       int
	GithubUsername string
	AvatarUrl      string
	Email          string
	IsAdmin        bool
}
