package goal

import "time"

type Goal struct {
	Id        int       `json:"id"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GoalContribution struct {
	Id                  int       `json:"id"`
	GoalId              int       `json:"goalId"`
	ContributionScoreId int       `json:"contributionScoreId"`
	TargetCount         int       `json:"targetCount"`
	IsCustom            bool      `json:"isCustom"`
	SetByUserId         int       `json:"setByUserId"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type CustomGoalLevelTarget struct {
	ContributionType string `json:"contributionType"`
	Target           int    `json:"target"`
}

type UserGoalLevelProgress struct {
	ContributionType string `json:"contributionType"`
	TargetCount      int    `json:"targetCount"`
	AchievedCount    int    `json:"achievedCount"`
}
