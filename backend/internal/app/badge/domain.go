package badge

import "time"

type Badge struct {
	Id        int
	UserId    int
	BadgeType string
	EarnedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
