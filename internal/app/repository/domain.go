package repository

import "time"

type RepoOWner struct {
	Login string `json:"login"`
}

type FetchRepositoryDetailsResponse struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	LanguagesURL    string    `json:"languages_url"`
	UpdateDate      time.Time `json:"updated_at"`
	RepoOwnerName   RepoOWner `json:"owner"`
	ContributorsUrl string    `json:"contributors_url"`
	RepoUrl         string    `json:"html_url"`
}

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

type FetchRepoContributorsResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"login"`
	AvatarUrl     string `json:"avatar_url"`
	GithubUrl     string `json:"html_url"`
	Contributions int    `json:"contributions"`
}

type FetchParticularRepoDetails struct {
	Repository
	Contributors []FetchRepoContributorsResponse
}

type Contribution struct {
	Id                  int
	UserId              int
	RepositoryId        int
	ContributionScoreId int
	ContributionType    string
	BalanceChange       int
	ContributedAt       time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type LanguagePercent struct {
	Name       string
	Bytes      int
	Percentage float64
}