package transaction

import "time"

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
