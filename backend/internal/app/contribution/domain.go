package contribution

import "time"

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

type ContributionScore struct {
	Id               int       `json:"id"`
	AdminId          int       `json:"adminId"`
	ContributionType string    `json:"contributionType"`
	Score            int       `json:"score"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
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

type MonthlyContributionSummary struct {
	Type       string    `json:"type"`
	Count      int       `json:"count"`
	TotalCoins int       `json:"totalCoins"`
	Month      time.Time `json:"month"`
}

type FetchUserContributionsResponse struct {
	Contribution
	Repository
}
