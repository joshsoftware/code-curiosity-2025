package user

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/goal"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	userRepository repository.UserRepository
	goalService    goal.GoalService
}

type Service interface {
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByGithubId(ctx context.Context, githubId int) (User, error)
	CreateUser(ctx context.Context, userInfo CreateUserRequestBody) (User, error)
	UpdateUserEmail(ctx context.Context, email string) error
	UpdateCurrentActiveGoalId(ctx context.Context, level string) (int, error)
}

func NewService(userRepository repository.UserRepository, goalService goal.GoalService) Service {
	return &service{
		userRepository: userRepository,
		goalService:    goalService,
	}
}

func (s *service) GetUserById(ctx context.Context, userId int) (User, error) {
	userInfo, err := s.userRepository.GetUserById(ctx, nil, userId)
	if err != nil {
		slog.Error("failed to get user by id", "error", err)
		return User{}, err
	}

	return User(userInfo), nil

}

func (s *service) GetUserByGithubId(ctx context.Context, githubId int) (User, error) {
	userInfo, err := s.userRepository.GetUserByGithubId(ctx, nil, githubId)
	if err != nil {
		slog.Error("failed to get user by github id", "error", err)
		return User{}, err
	}

	return User(userInfo), nil
}

func (s *service) CreateUser(ctx context.Context, userInfo CreateUserRequestBody) (User, error) {
	user, err := s.userRepository.CreateUser(ctx, nil, repository.CreateUserRequestBody(userInfo))
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return User{}, apperrors.ErrUserCreationFailed
	}

	return User(user), nil
}

func (s *service) UpdateUserEmail(ctx context.Context, email string) error {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return apperrors.ErrInternalServer
	}

	err := s.userRepository.UpdateUserEmail(ctx, nil, userId, email)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		return err
	}

	return nil
}

func (s *service) UpdateCurrentActiveGoalId(ctx context.Context, level string) (int, error) {
	userIdCtxVal := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdCtxVal.(int)

	if !ok {
		slog.Error("error obtaining user id from context")
		return 0, apperrors.ErrInternalServer
	}

	goalId, err := s.goalService.GetGoalIdByGoalLevel(ctx, level)

	if err != nil {
		slog.Error("[user service] error occured while fetching goal id by goal level")
		return 0, err
	}

	goalId, err = s.userRepository.UpdateCurrentActiveGoalId(ctx, nil, userId, goalId)

	if err != nil {
		slog.Error("[user service] Failed to update current active goal id", "error", err)
	}

	return goalId, err
}
