package badge

import "time"

type Badge struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	BadgeType string    `json:"badgeType"`
	EarnedAt  time.Time `json:"earnedAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
