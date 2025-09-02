package repository

import (
	"database/sql"
	"time"
)

type User struct {
	Id                  int           `db:"id"`
	GithubId            int           `db:"github_id"`
	GithubUsername      string        `db:"github_username"`
	Email               string        `db:"email"`
	AvatarUrl           string        `db:"avatar_url"`
	CurrentBalance      int           `db:"current_balance"`
	CurrentActiveGoalId sql.NullInt64 `db:"current_active_goal_id"`
	IsBlocked           bool          `db:"is_blocked"`
	IsAdmin             bool          `db:"is_admin"`
	Password            string        `db:"password"`
	IsDeleted           bool          `db:"is_deleted"`
	DeletedAt           sql.NullTime  `db:"deleted_at"`
	CreatedAt           time.Time     `db:"created_at"`
	UpdatedAt           time.Time     `db:"updated_at"`
}

type CreateUserRequestBody struct {
	GithubId       int    `db:"github_id"`
	GithubUsername string `db:"github_username"`
	AvatarUrl      string `db:"avatar_url"`
	Email          string `db:"email"`
	IsAdmin        bool   `db:"is_admin"`
	IsBlocked      bool   `db:"is_blocked"`
}

type Contribution struct {
	Id                  int       `db:"id"`
	UserId              int       `db:"user_id"`
	RepositoryId        int       `db:"repository_id"`
	ContributionScoreId int       `db:"contribution_score_id"`
	ContributionType    string    `db:"contribution_type"`
	BalanceChange       int       `db:"balance_change"`
	ContributedAt       time.Time `db:"contributed_at"`
	GithubEventId       string    `db:"github_event_id"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type Repository struct {
	Id              int       `db:"id"`
	GithubRepoId    int       `db:"github_repo_id"`
	RepoName        string    `db:"repo_name"`
	Description     string    `db:"description"`
	LanguagesUrl    string    `db:"languages_url"`
	RepoUrl         string    `db:"repo_url"`
	OwnerName       string    `db:"owner_name"`
	UpdateDate      time.Time `db:"update_date"`
	ContributorsUrl string    `db:"contributors_url"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type ContributionScore struct {
	Id               int       `db:"id"`
	AdminId          int       `db:"admin_id"`
	ContributionType string    `db:"contribution_type"`
	Score            int       `db:"score"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
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

type LeaderboardUser struct {
	Id             int    `db:"id"`
	GithubUsername string `db:"github_username"`
	AvatarUrl      string `db:"avatar_url"`
	CurrentBalance int    `db:"current_balance"`
	Rank           int    `db:"rank"`
}

type MonthlyContributionSummary struct {
	Type       string    `db:"contribution_type"`
	Count      int       `db:"contribution_count"`
	TotalCoins int       `db:"total_coins"`
	Month      time.Time `db:"month"`
}

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
	Id        int       `db:"id"`
	Level     string    `db:"level"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserGoal struct {
	Id             int       `db:"id"`
	UserId         int       `db:"user_id"`
	GoalLevelId    int       `db:"goal_level_id"`
	Status         string    `db:"status"`
	MonthStartedAt time.Time `db:"month_started_at"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type GoalLevelTarget struct {
	Id                  int       `db:"id"`
	GoalLevelId         int       `db:"goal_level_id"`
	ContributionScoreId int       `db:"contribution_score_id"`
	Target              int       `db:"target"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type UserGoalTarget struct {
	Id                  int       `db:"id"`
	UserGoalId          int       `db:"user_goal_id"`
	ContributionScoreId int       `db:"contribution_score_id"`
	Target              int       `db:"target"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type UserGoalProgress struct {
	UserGoalTargetId int `db:"user_goal_target_id"`
	ContributionId   int `db:"contribution_id"`
}

type Badge struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	BadgeType string    `db:"badge_type"`
	EarnedAt  time.Time `db:"earned_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AdminLoginRequest struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

type ConfigureContributionTypeScore struct {
	ContributionType string `db:"contribution_type"`
	Score            int    `db:"score"`
}

type GoalSummary struct {
	Id                   int       `db:"id"`
	UserId               int       `db:"user_id"`
	SnapshotDate         time.Time `db:"snapshot_date"`
	IncompleteGoalsCount int       `db:"incomplete_goals_count"`
	TargetSet            int       `db:"target_set"`
	TargetCompleted      int       `db:"target_completed"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
