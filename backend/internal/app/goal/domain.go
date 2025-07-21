package goal

import "time"

type Goal struct {
	Id        int
	Level     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GoalContribution struct {
	Id                  int
	GoalId              int
	ContributionScoreId int
	TargetCount         int
	IsCustom            bool
	SetByUserId         int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type CustomGoalLevelTarget struct {
	ContributionType string `json:"contribution_type"`
	Target           int    `json:"target"`
}
