package contribution

import "time"

type ContributionResponse struct {
	ID         string    `bigquery:"id"`
	Type       string    `bigquery:"type"`
	ActorID    int       `bigquery:"actor_id"`
	ActorLogin string    `bigquery:"actor_login"`
	RepoID     int       `bigquery:"repo_id"`
	RepoName   string    `bigquery:"repo_name"`
	RepoUrl    string    `bigquery:"repo_url"`
	Payload    string    `bigquery:"payload"`
	CreatedAt  time.Time `bigquery:"created_at"`
}

type Repository struct {
	Id           int
	GithubRepoId int
	RepoName     string
	Description  string
	LanguagesUrl string
	RepoUrl      string
	OwnerName    string
	UpdateDate   time.Time
	CreatedAt    int64
	UpdatedAt    int64
}

type Contribution struct {
	Id                  int
	UserId              int
	RepositoryId        int
	ContributionScoreId int
	ContributionType    string
	BalanceChange       int
	ContributedAt       time.Time
	CreatedAt           int64
	UpdatedAt           int64
}
