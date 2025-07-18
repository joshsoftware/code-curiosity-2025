package goal

import (
	"context"
	"log/slog"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/badge"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	goalRepository         repository.GoalRepository
	contributionRepository repository.ContributionRepository
	badgeService           badge.Service
}

type Service interface {
	ListGoalLevels(ctx context.Context) ([]Goal, error)
	GetGoalIdByGoalLevel(ctx context.Context, level string) (int, error)
	ListGoalLevelTargetDetail(ctx context.Context, userId int) ([]GoalContribution, error)
	CreateCustomGoalLevelTarget(ctx context.Context, userId int, customGoalLevelTarget []CustomGoalLevelTarget) ([]GoalContribution, error)
	ListGoalLevelAchievedTarget(ctx context.Context, userId int) (map[string]int, error)
}

func NewService(goalRepository repository.GoalRepository, contributionRepository repository.ContributionRepository, badgeService badge.Service) Service {
	return &service{
		goalRepository:         goalRepository,
		contributionRepository: contributionRepository,
		badgeService:           badgeService,
	}
}

func (s *service) ListGoalLevels(ctx context.Context) ([]Goal, error) {
	goals, err := s.goalRepository.ListGoalLevels(ctx, nil)
	if err != nil {
		slog.Error("error fetching goal levels", "error", err)
		return nil, err
	}

	serviceGoals := make([]Goal, len(goals))

	for i, g := range goals {
		serviceGoals[i] = Goal(g)
	}

	return serviceGoals, nil
}

func (s *service) GetGoalIdByGoalLevel(ctx context.Context, level string) (int, error) {
	goalId, err := s.goalRepository.GetGoalIdByGoalLevel(ctx, nil, level)

	if err != nil {
		slog.Error("failed to get goal id by goal level", "error", err)
		return 0, err
	}

	return goalId, err
}

func (s *service) ListGoalLevelTargetDetail(ctx context.Context, userId int) ([]GoalContribution, error) {
	goalLevelTargets, err := s.goalRepository.ListUserGoalLevelTargets(ctx, nil, userId)
	if err != nil {
		slog.Error("error fetching goal level targets", "error", err)
		return nil, err
	}

	serviceGoalLevelTargets := make([]GoalContribution, len(goalLevelTargets))
	for i, g := range goalLevelTargets {
		serviceGoalLevelTargets[i] = GoalContribution(g)
	}

	return serviceGoalLevelTargets, nil
}

func (s *service) CreateCustomGoalLevelTarget(ctx context.Context, userID int, customGoalLevelTarget []CustomGoalLevelTarget) ([]GoalContribution, error) {
	customGoalLevelId, err := s.GetGoalIdByGoalLevel(ctx, "Custom")
	if err != nil {
		slog.Error("error fetching custom goal level id", "error", err)
		return nil, err
	}
	var goalContributions []GoalContribution

	goalContributionInfo := make([]GoalContribution, len(customGoalLevelTarget))
	for i, c := range customGoalLevelTarget {
		goalContributionInfo[i].GoalId = customGoalLevelId

		contributionScoreDetails, err := s.contributionRepository.GetContributionScoreDetailsByContributionType(ctx, nil, c.ContributionType)
		if err != nil {
			slog.Error("error fetching contribution score details by type", "error", err)
			return nil, err
		}

		goalContributionInfo[i].ContributionScoreId = contributionScoreDetails.Id
		goalContributionInfo[i].TargetCount = c.Target
		goalContributionInfo[i].SetByUserId = userID

		goalContribution, err := s.goalRepository.CreateCustomGoalLevelTarget(ctx, nil, repository.GoalContribution(goalContributionInfo[i]))
		if err != nil {
			slog.Error("error creating custom goal level target", "error", err)
			return nil, err
		}

		goalContributions = append(goalContributions, GoalContribution(goalContribution))
	}

	return goalContributions, nil
}

func (s *service) ListGoalLevelAchievedTarget(ctx context.Context, userId int) (map[string]int, error) {
	goalLevelSetTargets, err := s.goalRepository.ListUserGoalLevelTargets(ctx, nil, userId)
	if err != nil {
		slog.Error("error fetching goal level targets", "error", err)
		return nil, err
	}

	contributionTypes := make([]CustomGoalLevelTarget, len(goalLevelSetTargets))
	for i, g := range goalLevelSetTargets {
		contributionTypes[i].ContributionType, err = s.contributionRepository.GetContributionTypeByContributionScoreId(ctx, nil, g.ContributionScoreId)
		if err != nil {
			slog.Error("error fetching contribution type by contribution score id", "error", err)
			return nil, err
		}

		contributionTypes[i].Target = g.TargetCount
	}

	year := int(time.Now().Year())
	month := int(time.Now().Month())
	monthlyContributionCount, err := s.contributionRepository.ListMonthlyContributionSummary(ctx, nil, year, month, userId)
	if err != nil {
		slog.Error("error fetching monthly contribution count", "error", err)
		return nil, err
	}

	contributionsAchievedTarget := make(map[string]int, len(monthlyContributionCount))

	for _, m := range monthlyContributionCount {
		contributionsAchievedTarget[m.Type] = m.Count
	}

	var completedTarget int
	for _, c := range contributionTypes {
		if c.Target == contributionsAchievedTarget[c.ContributionType] {
			completedTarget += 1
		}
	}

	if completedTarget == len(goalLevelSetTargets) {
		userGoalLevel, err := s.goalRepository.GetUserActiveGoalLevel(ctx, nil, userId)
		if err != nil {
			slog.Error("error fetching user active gaol level", "error", err)
			return nil, err
		}

		_, err = s.badgeService.HandleBadgeCreation(ctx, userId, userGoalLevel)
		if err != nil {
			slog.Error("error handling user badge creation", "error", err)
			return nil, err
		}
	}

	return contributionsAchievedTarget, nil
}
