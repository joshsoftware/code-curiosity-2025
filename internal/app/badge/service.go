package badge

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type badgeService struct {
	badgeRepository repository.BadgeRepository
}

type BadgeService interface {
	GetBadgeDetailsOfUser(ctx context.Context) ([]Badge, error)
}

func NewBadgeService(badgeRepository repository.BadgeRepository) BadgeService {
	return &badgeService{
		badgeRepository: badgeRepository,
	}
}

func (bs *badgeService) GetBadgeDetailsOfUser(ctx context.Context) ([]Badge, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("(service) error obtaining user id from context")
		return nil, apperrors.ErrInternalServer
	}

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
