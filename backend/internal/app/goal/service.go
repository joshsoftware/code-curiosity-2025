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
	GetUserActiveGoalLevel(ctx context.Context, userId int) (string, error)
	CreateCustomGoalLevelTarget(ctx context.Context, userId int, customGoalLevelTarget []CustomGoalLevelTarget) ([]GoalContribution, error)
	ListUserGoalLevelProgress(ctx context.Context, userId int) ([]UserGoalLevelProgress, error)
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

func (s *service) GetUserActiveGoalLevel(ctx context.Context, userId int) (string, error) {
	userGoalLevel, err := s.goalRepository.GetUserActiveGoalLevel(ctx, nil, userId)
	if err != nil {
		slog.Error("error fetching user active gaol level", "error", err)
		return "", err
	}

	return userGoalLevel, nil
}

func (s *service) CreateCustomGoalLevelTarget(ctx context.Context, userId int, customGoalLevelTarget []CustomGoalLevelTarget) ([]GoalContribution, error) {
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
		goalContributionInfo[i].SetByUserId = userId

		goalContribution, err := s.goalRepository.CreateCustomGoalLevelTarget(ctx, nil, repository.GoalContribution(goalContributionInfo[i]))
		if err != nil {
			slog.Error("error creating custom goal level target", "error", err)
			return nil, err
		}

		goalContributions = append(goalContributions, GoalContribution(goalContribution))
	}

	return goalContributions, nil
}


func (s *service) ListUserGoalLevelProgress(ctx context.Context, userId int) ([]UserGoalLevelProgress, error) {
	goalLevelSetTargets, err := s.goalRepository.ListUserGoalLevelTargets(ctx, nil, userId)
	if err != nil {
		slog.Error("error fetching goal level targets", "error", err)
		return nil, err
	}

	year := int(time.Now().Year())
	month := int(time.Now().Month())
	monthlyContributionCount, err := s.contributionRepository.ListMonthlyContributionSummary(ctx, nil, year, month, userId)
	if err != nil {
		slog.Error("error fetching monthly contribution count", "error", err)
		return nil, err
	}

	contributionCountMap := make(map[string]int)
	for _, m := range monthlyContributionCount {
		contributionCountMap[m.Type] = m.Count
	}

	userGoalLevelProgress := make([]UserGoalLevelProgress, len(goalLevelSetTargets))
	var contributionsCompleted int

	for i, g := range goalLevelSetTargets {
		contributionType, err := s.contributionRepository.GetContributionTypeByContributionScoreId(ctx, nil, g.ContributionScoreId)
		if err != nil {
			slog.Error("error")
			return nil, err
		}

		userGoalLevelProgress[i].ContributionType = contributionType
		userGoalLevelProgress[i].TargetCount = g.TargetCount
		userGoalLevelProgress[i].AchievedCount = contributionCountMap[contributionType]

		if userGoalLevelProgress[i].AchievedCount == g.TargetCount {
			contributionsCompleted++
		}

		if contributionsCompleted == len(goalLevelSetTargets) {
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
	}

	return userGoalLevelProgress, nil
}
