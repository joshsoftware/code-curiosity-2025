package goal

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/badge"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	goalRepository         repository.GoalRepository
	contributionRepository repository.ContributionRepository
	badgeService           badge.Service
}

type Service interface {
	ListGoalLevels(ctx context.Context) ([]GoalLevel, error)
	CreateUserGoalInProgress(ctx context.Context, userSelecetdGoal CreateUserGoalRequest, userId int) (UserGoal, error)
	CreateCustomUserGoalTarget(ctx context.Context, userSelectedCustomGoals []CustomTargetRequest, createdUserGoal UserGoal) ([]UserGoalTarget, error)
	SyncUserGoalProgress(ctx context.Context, userGoalTargets []UserGoalTarget, monthStartedAt time.Time, userId int) ([]UserGoalProgress, error)
	ResetUserCurrentGoalStatus(ctx context.Context, userId int) (UserGoal, error)
	GetUserCurrentGoalStatus(ctx context.Context, userId int) (*GetUserCurrentGoalStatusResponse, error)
	AllocateBadge(ctx context.Context, userId int) error
	UpdateUserGoalStatusMonthly(ctx context.Context) error
	SyncUserGoalProgressWithContributions(ctx context.Context, userId int) error
	CreateUserGoalSummary(ctx context.Context, userId int) (GoalSummary, error)
	FetchUserGoalSummary(ctx context.Context, userId int) ([]GoalSummary, error)
}

func NewService(goalRepository repository.GoalRepository, contributionRepository repository.ContributionRepository, badgeService badge.Service) Service {
	return &service{
		goalRepository:         goalRepository,
		contributionRepository: contributionRepository,
		badgeService:           badgeService,
	}
}

func (s *service) ListGoalLevels(ctx context.Context) ([]GoalLevel, error) {
	goals, err := s.goalRepository.ListGoalLevels(ctx, nil)
	if err != nil {
		slog.Error("error fetching goal levels", "error", err)
		return nil, err
	}

	serviceGoals := make([]GoalLevel, len(goals))
	for i, g := range goals {
		serviceGoals[i] = GoalLevel(g)
	}

	return serviceGoals, nil
}

func (s *service) CreateUserGoalInProgress(ctx context.Context, userSelecetdGoal CreateUserGoalRequest, userId int) (UserGoal, error) {
	now := time.Now().UTC()
	monthStartedAt := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	userCurrentGoal, err := s.goalRepository.GetUserCurrentGoal(ctx, nil, userId)
	if err == nil {
		slog.Error("user already has existing goal set for current month")
		return UserGoal(userCurrentGoal), apperrors.ErrUserGoalExists
	} else if !errors.Is(err, apperrors.ErrUserGoalNotFound) {
		slog.Error("error getting user goal for current month")
		return UserGoal{}, err
	}

	goalLevel, err := s.goalRepository.GetGoalLevelByLevel(ctx, nil, userSelecetdGoal.Level)
	if err != nil {
		slog.Error("error fetching goal id by goal level", "error", err)
		return UserGoal{}, err
	}

	userGoal := UserGoal{
		UserId:         userId,
		GoalLevelId:    goalLevel.Id,
		Status:         GoalStatusInProgress,
		MonthStartedAt: monthStartedAt,
	}

	createdUserGoal, err := s.goalRepository.CreateUserGoalInProgress(ctx, nil, repository.UserGoal(userGoal))
	if err != nil {
		slog.Error("failed to create user goal status", "error", err)
		return UserGoal{}, err
	}

	var createdUserGoalTargets []UserGoalTarget

	//check if custom
	if userSelecetdGoal.Level == GoalLevelCustom {
		createdUserGoalTargets, err = s.CreateCustomUserGoalTarget(ctx, userSelecetdGoal.CustomTargets, UserGoal(createdUserGoal))
		if err != nil {
			slog.Error("error creating custom goal target", "error", err)
			return UserGoal{}, err
		}
	}

	//check if goal level is not custom
	if userSelecetdGoal.Level != GoalLevelCustom {
		goalLevelTargets, err := s.goalRepository.FetchGoalLevelTargetByGoalLevel(ctx, nil, goalLevel)
		if err != nil {
			slog.Error("error fetching goal level target", "error", err)
			return UserGoal{}, err
		}

		for _, g := range goalLevelTargets {
			userGoalTarget := UserGoalTarget{
				UserGoalId:          createdUserGoal.Id,
				ContributionScoreId: g.ContributionScoreId,
				Target:              g.Target,
			}

			createdUserGoalTarget, err := s.goalRepository.CreateUserGoalTarget(ctx, nil, repository.UserGoalTarget(userGoalTarget))
			if err != nil {
				slog.Error("error creeating user goal target", "error", err)
				return UserGoal{}, err
			}

			createdUserGoalTargets = append(createdUserGoalTargets, UserGoalTarget(createdUserGoalTarget))
		}
	}

	_, err = s.SyncUserGoalProgress(ctx, createdUserGoalTargets, monthStartedAt, userId)
	if err != nil {
		slog.Error("error syncing user goal progress", "error", err)
		return UserGoal{}, err
	}

	return UserGoal(createdUserGoal), nil
}

func (s *service) CreateCustomUserGoalTarget(ctx context.Context, userSelectedCustomGoals []CustomTargetRequest, createdUserGoal UserGoal) ([]UserGoalTarget, error) {
	createdUserGoalTargets := make([]UserGoalTarget, len(userSelectedCustomGoals))

	for _, userSelectedCustomGoal := range userSelectedCustomGoals {
		contributionScoreDetails, err := s.contributionRepository.GetContributionScoreDetailsByContributionType(ctx, nil, userSelectedCustomGoal.ContributionType)
		if err != nil {
			slog.Error("error getting contirbution score details for given contribution type", "error", err)
			return nil, err
		}

		userGoalTarget := UserGoalTarget{
			UserGoalId:          createdUserGoal.Id,
			ContributionScoreId: contributionScoreDetails.Id,
			Target:              userSelectedCustomGoal.Target,
		}

		createdUserGoalTarget, err := s.goalRepository.CreateUserGoalTarget(ctx, nil, repository.UserGoalTarget(userGoalTarget))
		if err != nil {
			slog.Error("error creeating user goal target", "error", err)
			return nil, err
		}

		createdUserGoalTargets = append(createdUserGoalTargets, UserGoalTarget(createdUserGoalTarget))
	}

	return createdUserGoalTargets, nil
}

func (s *service) SyncUserGoalProgress(ctx context.Context, userGoalTargets []UserGoalTarget, monthStartedAt time.Time, userId int) ([]UserGoalProgress, error) {
	userContributionsForMonth, err := s.contributionRepository.FetchUserContributionsForMonth(ctx, nil, userId, monthStartedAt)
	if err != nil {
		slog.Error("error fetching user contributions for month", "error", err)
		return nil, err
	}

	contributionMap := make(map[int][]Contribution)
	for _, c := range userContributionsForMonth {
		contributionMap[c.ContributionScoreId] = append(contributionMap[c.ContributionScoreId], Contribution(c))
	}

	var createdUserGoalProgresses []UserGoalProgress

	for _, target := range userGoalTargets {
		if contributions, ok := contributionMap[target.ContributionScoreId]; ok {
			for _, contribution := range contributions {
				userGoalProgress := UserGoalProgress{
					UserGoalTargetId: target.Id,
					ContributionId:   contribution.Id,
				}

				created, err := s.goalRepository.CreateUserGoalProgress(ctx, nil, repository.UserGoalProgress(userGoalProgress))
				if err != nil {
					slog.Error("error creating user goal progress", "error", err)
					continue
				}

				createdUserGoalProgresses = append(createdUserGoalProgresses, UserGoalProgress(created))
			}
		}
	}

	return createdUserGoalProgresses, nil
}

func (s *service) ResetUserCurrentGoalStatus(ctx context.Context, userId int) (UserGoal, error) {
	userCurrentGoal, err := s.goalRepository.GetUserCurrentGoal(ctx, nil, userId)
	if err != nil {
		slog.Error("error getting user goal for current month")
		return UserGoal{}, err
	}

	if time.Since(userCurrentGoal.CreatedAt) > 48*time.Hour || userCurrentGoal.Status != GoalStatusInProgress {
		slog.Error("cannot reset goal", "error", err)
		return UserGoal{}, apperrors.ErrFailedResettingGoal
	}

	userGoal := UserGoal{
		Id:     userCurrentGoal.Id,
		Status: GoalStatusIncomplete,
	}
	updatedUserGoal, err := s.goalRepository.UpdateUserGoalStatus(ctx, nil, repository.UserGoal(userGoal))
	if err != nil {
		slog.Error("error updating goal status for user", "error", err)
		return UserGoal{}, err
	}

	return UserGoal(updatedUserGoal), nil
}

func (s *service) GetUserCurrentGoalStatus(ctx context.Context, userId int) (*GetUserCurrentGoalStatusResponse, error) {
	userCurrentGoal, err := s.goalRepository.GetUserCurrentGoal(ctx, nil, userId)
	if err != nil {
		slog.Error("error getting user goal for current month")
		return nil, err
	}

	goalLevel, err := s.goalRepository.GetGoalLevelById(ctx, nil, userCurrentGoal.GoalLevelId)
	if err != nil {
		slog.Error("error fetching goal leve by goal level id", "error", err)
		return nil, err
	}

	userCurrentGoalTargets, err := s.goalRepository.ListUserGoalTargetsByUserGoalId(ctx, nil, userCurrentGoal.Id)
	if err != nil {
		slog.Error("error fetching user goal targets by user goal id", "error", err)
		return nil, err
	}

	goalTargetProgresses := make([]UserGoalTargetProgress, 0, len(userCurrentGoalTargets))

	var totalTargetsCompleted int
	totalTargets := len(userCurrentGoalTargets)

	for _, userCurrentGoalTarget := range userCurrentGoalTargets {

		contributionType, err := s.contributionRepository.GetContributionTypeByContributionScoreId(ctx, nil, userCurrentGoalTarget.ContributionScoreId)
		if err != nil {
			slog.Error("error fetching contribution type by contribution score id", "error", err)
			return nil, err
		}

		contributionProgressCount, err := s.goalRepository.GetContributionProgressCount(ctx, nil, userCurrentGoalTarget.Id)
		if err != nil {
			slog.Error("error fetching contribution progress count", "error", err)
			return nil, err
		}

		goalTargetProgress := UserGoalTargetProgress{
			ContributionType: contributionType,
			Target:           userCurrentGoalTarget.Target,
			Progress:         contributionProgressCount,
		}

		if goalTargetProgress.Target <= goalTargetProgress.Progress {
			totalTargetsCompleted++
		}

		goalTargetProgresses = append(goalTargetProgresses, goalTargetProgress)
	}

	userCurrentGoalStatusResponse := GetUserCurrentGoalStatusResponse{
		UserGoalId:         userCurrentGoal.Id,
		Level:              goalLevel.Level,
		Status:             userCurrentGoal.Status,
		MonthStartedAt:     userCurrentGoal.MonthStartedAt,
		CreatedAt:          userCurrentGoal.UpdatedAt,
		UpdatedAt:          userCurrentGoal.UpdatedAt,
		GoalTargetProgress: goalTargetProgresses,
	}

	if totalTargets <= totalTargetsCompleted {
		userGoal := UserGoal{
			Id:     userCurrentGoal.Id,
			Status: GoalStatusCompleted,
		}
		_, err = s.goalRepository.UpdateUserGoalStatus(ctx, nil, repository.UserGoal(userGoal))
		if err != nil {
			slog.Error("error updating user goal status to complete", "error", err)
			return nil, err
		}

		_, err = s.badgeService.HandleBadgeCreation(ctx, userId, goalLevel.Level)
		if err != nil {
			slog.Error("error creating badge", "error", err)
			return nil, err
		}
	}

	return &userCurrentGoalStatusResponse, nil
}

func (s *service) SyncUserGoalProgressWithContributions(ctx context.Context, userId int) error {
	userCurrentGoal, err := s.goalRepository.GetUserCurrentGoal(ctx, nil, userId)
	if err != nil {
		slog.Error("error getting user goal for current month", "error", err)
		return err
	}

	userCurrentGoalTargets, err := s.goalRepository.ListUserGoalTargetsByUserGoalId(ctx, nil, userCurrentGoal.Id)
	if err != nil {
		slog.Error("error fetching user goal targets by user goal id", "error", err)
		return err
	}

	serviceUserCurrentGoalTargets := make([]UserGoalTarget, len(userCurrentGoalTargets))
	for i, userCurrentGoalTarget := range userCurrentGoalTargets {
		serviceUserCurrentGoalTargets[i] = UserGoalTarget(userCurrentGoalTarget)
	}

	now := time.Now().UTC()
	monthStartedAt := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	_, err = s.SyncUserGoalProgress(ctx, serviceUserCurrentGoalTargets, monthStartedAt, userId)
	if err != nil {
		slog.Error("error syncing user goal progress with contributions", "error", err)
		return err
	}

	return nil
}

func (s *service) AllocateBadge(ctx context.Context, userId int) error {
	userCurrentGoalStatus, err := s.GetUserCurrentGoalStatus(ctx, userId)
	if err != nil {
		slog.Error("error fetching user current goal status", "error", err)
		return err
	}

	var totalTargetsCompleted int
	totalTargets := len(userCurrentGoalStatus.GoalTargetProgress)
	for _, goalTargetProgress := range userCurrentGoalStatus.GoalTargetProgress {
		if goalTargetProgress.Target <= goalTargetProgress.Progress {
			totalTargetsCompleted++
		}
	}

	if totalTargets <= totalTargetsCompleted {
		userGoal := UserGoal{
			Id:     userCurrentGoalStatus.UserGoalId,
			Status: GoalStatusCompleted,
		}
		_, err = s.goalRepository.UpdateUserGoalStatus(ctx, nil, repository.UserGoal(userGoal))
		if err != nil {
			slog.Error("error updating user goal status to complete", "error", err)
			return err
		}

		_, err = s.badgeService.HandleBadgeCreation(ctx, userId, userCurrentGoalStatus.Level)
		if err != nil {
			slog.Error("error creating badge", "error", err)
			return err
		}
	}

	return nil
}

func (s *service) UpdateUserGoalStatusMonthly(ctx context.Context) error {
	userGoals, err := s.goalRepository.FetchInProgressUserGoalsOfPreviousMonth(ctx, nil)
	if err != nil {
		slog.Error("error getting user goals for previous month", "error", err)
		return err
	}

	for _, userGoal := range userGoals {
		userUpdatedGoal := UserGoal{
			Id:     userGoal.Id,
			Status: GoalStatusCompleted,
		}
		_, err = s.goalRepository.UpdateUserGoalStatus(ctx, nil, repository.UserGoal(userUpdatedGoal))
		if err != nil {
			slog.Error("error updating users goal for previous month", "error", err)
			return err
		}
	}

	return nil
}

func (s *service) CreateUserGoalSummary(ctx context.Context, userId int) (GoalSummary, error) {
	userIncompleteGoalCount, err := s.goalRepository.CalculateUserIncompleteGoalsUntilDay(ctx, nil, userId)
	if err != nil {
		slog.Error("error calculating user incomplete goalstatus until day", "error", err)
		return GoalSummary{}, err
	}

	userCurrentGoalStatus, err := s.GetUserCurrentGoalStatus(ctx, userId)
	if err != nil {
		slog.Error("error getting user current goal status", "error", err)
		return GoalSummary{}, err
	}

	var totalTargetSet int
	var totalTargetCompleted int
	for _, s := range userCurrentGoalStatus.GoalTargetProgress {
		totalTargetSet += s.Target
		totalTargetCompleted += s.Progress
	}

	userMonthlyGoalSummary := GoalSummary{
		UserId:               userId,
		SnapshotDate:         time.Now().UTC(),
		IncompleteGoalsCount: userIncompleteGoalCount,
		TargetSet:            totalTargetSet,
		TargetCompleted:      totalTargetCompleted,
	}

	createdUserGoalSummary, err := s.goalRepository.CreateUserGoalSummary(ctx, nil, repository.GoalSummary(userMonthlyGoalSummary))
	if err != nil {
		slog.Error("error creating user goal summary", "error", err)
		return GoalSummary{}, err
	}

	return GoalSummary(createdUserGoalSummary), nil
}

func (s *service) FetchUserGoalSummary(ctx context.Context, userId int) ([]GoalSummary, error) {
	usersGoalSummary, err := s.goalRepository.FetchUserGoalSummary(ctx, nil, userId)
	if err != nil {
		slog.Error("error fetching user goal summary", "error", err)
		return nil, err
	}

	serviceUserGoalSummary := make([]GoalSummary, len(usersGoalSummary))
	for i, userGoalSummary := range usersGoalSummary {
		serviceUserGoalSummary[i] = GoalSummary(userGoalSummary)
	}

	return serviceUserGoalSummary, nil
}
