package transaction

import "time"

type Transaction struct {
	Id                int
	UserId            int
	ContributionId    int
	IsRedeemed        bool
	IsGained          bool
	TransactedBalance int
	TransactedAt      time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
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
