package goal

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type goalService struct {
	goalRepository repository.GoalRepository
}

type GoalService interface {
	GetGoalIdByGoalLevel(ctx context.Context, level string) (int, error)
}

func NewGoalService(goalRepository repository.GoalRepository) GoalService {
	return &goalService{
		goalRepository: goalRepository,
	}
}

func (gs *goalService) GetGoalIdByGoalLevel(ctx context.Context, level string) (int, error) {
	goalId, err := gs.goalRepository.GetGoalIdByGoalLevel(ctx, nil, level)

	if err != nil {
		slog.Error("[Goal service] Failed to get goal id by goal level", "error", err)
		return 0, err
	}

	return goalId, err
}
