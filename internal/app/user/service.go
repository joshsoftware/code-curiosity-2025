package user

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	userRepository repository.UserRepository
}

type Service interface {
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByGithubId(ctx context.Context, githubId int) (User, error)
	CreateUser(ctx context.Context, userInfo CreateUserRequestBody) (User, error)
	UpdateUserEmail(ctx context.Context, email string) error
	UpdateUserCurrentBalance(ctx context.Context, transaction Transaction) error
}

func NewService(userRepository repository.UserRepository) Service {
	return &service{
		userRepository: userRepository,
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

func (s *service) UpdateUserCurrentBalance(ctx context.Context, transaction Transaction) error {
	user, err := s.GetUserById(ctx, transaction.UserId)
	if err != nil {
		slog.Error("error obtaining user by id", "error", err)
		return err
	}

	user.CurrentBalance += transaction.TransactedBalance

	tx, ok := middleware.ExtractTxFromContext(ctx)
	if !ok {
		slog.Error("error obtaining tx from context")
	}

	err = s.userRepository.UpdateUserCurrentBalance(ctx, tx, repository.User(user))
	if err != nil {
		slog.Error("error updating user current balance", "error", err)
		return err
	}

	return nil
}
