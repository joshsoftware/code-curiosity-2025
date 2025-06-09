package badge

import (
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type Badge struct {
	BadgeId   int       `json:"badge_id"`
	BadgeType string    `json:"badge_type"`
	EarnedAt  time.Time `json:"earned_at"`
}

func FromRepositoryBadge(repoBadge repository.Badge) Badge {
	return Badge(repoBadge);
}
