package badge

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	badgeRepository repository.BadgeRepository
}

type Service interface {
	HandleBadgeCreation(ctx context.Context, userId int, badgeType string) (Badge, error)
	GetBadgeDetailsOfUser(ctx context.Context, userId int) ([]Badge, error)
}

func NewService(badgeRepository repository.BadgeRepository) Service {
	return &service{
		badgeRepository: badgeRepository,
	}
}

func (s *service) HandleBadgeCreation(ctx context.Context, userId int, badgeType string) (Badge, error) {
	badge, err := s.badgeRepository.GetUserCurrentMonthBadge(ctx, nil, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			badge, err = s.badgeRepository.CreateBadge(ctx, nil, userId, badgeType)
			if err != nil {
				slog.Error("error creating badge for user", "error", err)
				return Badge{}, err
			}
		}
		slog.Error("error fetching current month badge for user", "error", err)
		return Badge{}, err
	}

	return Badge(badge), nil
}

func (s *service) GetBadgeDetailsOfUser(ctx context.Context, userId int) ([]Badge, error) {
	badges, err := s.badgeRepository.GetBadgeDetailsOfUser(ctx, nil, userId)

	if err != nil {
		slog.Error("(service) Failed to get the badge details", "error", err)
		return nil, err
	}

	serviceBadge := make([]Badge, len(badges))

	for i, b := range badges {
		serviceBadge[i] = Badge(b)
	}

	return serviceBadge, nil
}
