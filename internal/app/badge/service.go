package badge

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type badgeService struct {
	badgeRepository repository.BadgeRepository
}

type BadgeService interface {
	GetBadgeDetailsOfUser(ctx context.Context, userId int) ([]Badge, error)
}

func NewBadgeService(badgeRepository repository.BadgeRepository) BadgeService {
	return &badgeService{
		badgeRepository: badgeRepository,
	}
}

func (bs *badgeService) GetBadgeDetailsOfUser(ctx context.Context, userId int) ([]Badge, error) {
	badges, err := bs.badgeRepository.GetBadgeDetailsOfUser(ctx, nil, userId)

	if err != nil {
		slog.Error("(service) Failed to get the badge details", "error", err)
		return nil, err
	}

	finalBadges := make([]Badge, len(badges))

	for i, badge := range badges {
		finalBadges[i] = FromRepositoryBadge(badge)
	}

	return finalBadges, nil
}
