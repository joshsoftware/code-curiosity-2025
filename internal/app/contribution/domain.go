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

type ContributionScore struct {
	Id               int
	AdminId          int
	ContributionType string
	Score            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Transaction struct {
	Id                int       `db:"id"`
	UserId            int       `db:"user_id"`
	ContributionId    int       `db:"contribution_id"`
	IsRedeemed        bool      `db:"is_redeemed"`
	IsGained          bool      `db:"is_gained"`
	TransactedBalance int       `db:"transacted_balance"`
	TransactedAt      time.Time `db:"transacted_at"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type ContributionTypeSummary struct {
	ContributionType  string    `db:"contribution_type"`
	ContributionCount int       `db:"contribution_count"`
	TotalCoins        int       `db:"total_coins"`
	Month             time.Time `db:"month"`
}
