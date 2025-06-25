package goal

type GoalAction struct {
	ActionName        string `json:"name"`
	ActionTargetCount int    `json:"target_count"`
}

type GoalLevel struct {
	LevelName string       `json:"level_name"`
	Actions   []GoalAction `json:"actions,omitempty"`
}

type GoalId struct {
	GoalId int `json:"goal_id"`
}
