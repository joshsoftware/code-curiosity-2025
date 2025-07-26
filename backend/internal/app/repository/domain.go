package repository

import "time"

type Repository struct {
	Id              int       `json:"id"`
	GithubRepoId    int       `json:"githubRepoId"`
	RepoName        string    `json:"repoName"`
	Description     string    `json:"description"`
	LanguagesUrl    string    `json:"languagesUrl"`
	RepoUrl         string    `json:"repoUrl"`
	OwnerName       string    `json:"ownerName"`
	UpdateDate      time.Time `json:"updateDate"`
	ContributorsUrl string    `json:"contributorsUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type RepoLanguages map[string]int

type FetchUsersContributedReposResponse struct {
	Repository
	Languages        []string `json:"languages"`
	TotalCoinsEarned int      `json:"totalCoinsEarned"`
}

type ContributionResponse struct {
	ID         string    `bigquery:"id" json:"id"`
	Type       string    `bigquery:"type" json:"type"`
	ActorID    int       `bigquery:"actor_id" json:"actorId"`
	ActorLogin string    `bigquery:"actor_login" json:"actorLogin"`
	RepoID     int       `bigquery:"repo_id" json:"repoId"`
	RepoName   string    `bigquery:"repo_name" json:"repoName"`
	RepoUrl    string    `bigquery:"repo_url" json:"repoUrl"`
	Payload    string    `bigquery:"payload" json:"payload"`
	CreatedAt  time.Time `bigquery:"created_at" json:"createdAt"`
}

type Contribution struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"userId"`
	RepositoryId        int       `json:"repositoryId"`
	ContributionScoreId int       `json:"contributionScoreId"`
	ContributionType    string    `json:"contributionType"`
	BalanceChange       int       `json:"balanceChange"`
	ContributedAt       time.Time `json:"contributedAt"`
	GithubEventId       string    `json:"githubEventId"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type LanguagePercent struct {
	Name       string  `json:"name"`
	Bytes      int     `json:"bytes"`
	Percentage float64 `json:"percentage"`
}
