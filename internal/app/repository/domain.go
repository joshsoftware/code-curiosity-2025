package repository

import "time"

type Repository struct {
	Id              int
	GithubRepoId    int
	RepoName        string
	Description     string
	LanguagesUrl    string
	RepoUrl         string
	OwnerName       string
	UpdateDate      time.Time
	ContributorsUrl string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type RepoLanguages map[string]int

type FetchUsersContributedReposResponse struct {
	Repository
	Languages        []string
	TotalCoinsEarned int
}

type Contribution struct {
	Id                  int
	UserId              int
	RepositoryId        int
	ContributionScoreId int
	ContributionType    string
	BalanceChange       int
	ContributedAt       time.Time
	GithubEventId       string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type LanguagePercent struct {
	Name       string
	Bytes      int
	Percentage float64
}
