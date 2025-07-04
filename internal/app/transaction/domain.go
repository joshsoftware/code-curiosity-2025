package transaction

import "time"

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
