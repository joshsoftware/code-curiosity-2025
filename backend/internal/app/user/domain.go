package user

import (
	"database/sql"
	"time"
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

type CreateUserRequestBody struct {
	GithubId       int    `json:"githubId"`
	GithubUsername string `json:"githubUsername"`
	AvatarUrl      string `json:"avatarUrl"`
	Email          string `json:"email"`
	IsAdmin        bool   `json:"isAdmin"`
}

type Email struct {
	Email string `json:"email"`
}

type Transaction struct {
	Id                int       `json:"id"`
	UserId            int       `json:"userId"`
	ContributionId    int       `json:"contributionId"`
	IsRedeemed        bool      `json:"isRedeemed"`
	IsGained          bool      `json:"isGained"`
	TransactedBalance int       `json:"transactedBalance"`
	TransactedAt      time.Time `json:"transactedAt"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type LeaderboardUser struct {
	Id                    int    `json:"id"`
	GithubUsername        string `json:"githubUsername"`
	AvatarUrl             string `json:"avatarUrl"`
	ContributedReposCount int    `json:"contributedReposCount"`
	CurrentBalance        int    `json:"currentBalance"`
	Rank                  int    `json:"rank"`
}

type GoalLevel struct {
	Level string `json:"level"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BlockOrUnblockUserRequest struct {
	Block bool `json:"block"`
}
