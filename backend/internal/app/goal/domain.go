package goal

import "time"

const (
	GoalStatusInProgress = "inProgress"
	GoalStatusCompleted  = "completed"
	GoalStatusIncomplete = "incomplete"
)

const (
	GoalLevelBeginner     = "Beginner"
	GoalLevelIntermediate = "Intermediate"
	GoalLevelAdvanced     = "Advanced"
	GoalLevelCustom       = "Custom"
)

type GoalLevel struct {
	Id        int       `json:"id"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GoalLevelName struct {
	Level string `json:"level"`
}

type UserGoal struct {
	Id             int       `json:"id"`
	UserId         int       `json:"userId"`
	GoalLevelId    int       `json:"goalLevelId"`
	Status         string    `json:"status"`
	MonthStartedAt time.Time `json:"monthStartedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type UserGoalTarget struct {
	Id                  int       `json:"id"`
	UserGoalId          int       `json:"userGoalId"`
	ContributionScoreId int       `json:"contributionScoreId"`
	Target              int       `json:"target"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type UserGoalProgress struct {
	UserGoalTargetId int `json:"userGoalTargetId"`
	ContributionId   int `json:"contributionId"`
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

type CreateUserGoalRequest struct {
	Level         string                `json:"level"`
	CustomTargets []CustomTargetRequest `json:"customTargets,omitempty"`
}

type CustomTargetRequest struct {
	ContributionType string `json:"contributionType"`
	Target           int    `json:"target"`
}

type GetUserCurrentGoalStatusResponse struct {
	UserGoalId         int                      `json:"userGoalId"`
	Level              string                   `json:"level"`
	Status             string                   `json:"status"`
	MonthStartedAt     time.Time                `json:"monthStartedAt"`
	CreatedAt          time.Time                `json:"createdAt"`
	UpdatedAt          time.Time                `json:"updatedAt"`
	GoalTargetProgress []UserGoalTargetProgress `json:"goalTargetProgress"`
}

type UserGoalTargetProgress struct {
	ContributionType string `json:"contributionType"`
	Target           int    `json:"target"`
	Progress         int    `json:"progress"`
}

type UserGoalIdRequest struct {
	UserGoalId int `json:"userGoalId"`
}

type GoalSummary struct {
	Id                   int       `json:"id"`
	UserId               int       `json:"userId"`
	SnapshotDate         time.Time `json:"snapshotDate"`
	IncompleteGoalsCount int       `json:"incompleteGoalsCount"`
	TargetSet            int       `json:"targetSet"`
	TargetCompleted      int       `json:"targetCompleted"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

// type MonthlyGoalSummary struct {
// 	Day                  time.Time `json:"Day"`
// 	IncompleteGoalsCount int `json:"IncompleteGoalsCount"`
// 	TargetSet            int       `json:"TargetSet"`
// 	TargetCompleted      int       `json:"TargetCompleted"`
// }
