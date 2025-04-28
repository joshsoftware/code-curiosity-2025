package user

import (
	"context"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/constants"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	userRepository repository.UserRepository
}

type Service interface {
	GetUserByGithubId(ctx context.Context, githubId int) (User, error)
	CreateUser(ctx context.Context, userInfo CreateUserRequestBody) (User, error)
	GetLoggedInUser(ctx context.Context) (User, error)
	UpdateUserEmail(ctx context.Context, email string) error
}

func NewService(userRepository repository.UserRepository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) GetUserByGithubId(ctx context.Context, githubId int) (User, error) {
	tx, err := s.userRepository.BeginTx(ctx)
	if err != nil {
		slog.Error("failed to start user creation", "error", err)
		return User{}, err
	}

	defer func() {
		if txErr := s.userRepository.HandleTransaction(ctx, tx, err); txErr != nil {
			slog.Error("failed to handle transaction", "error", txErr)
			err = txErr
		}
	}()

	userInfo, err := s.userRepository.GetUserByGithubId(ctx, tx, githubId)
	if err != nil {
		slog.Error("failed to get user by id", "error", err)
		return User{}, err
	}

	return User(userInfo), nil
}

func (s *service) CreateUser(ctx context.Context, userInfo CreateUserRequestBody) (User, error) {
	tx, err := s.userRepository.BeginTx(ctx)
	if err != nil {
		slog.Error("failed to start user creation", "error", err)
		return User{}, err
	}

	defer func() {
		if txErr := s.userRepository.HandleTransaction(ctx, tx, err); txErr != nil {
			slog.Error("failed to handle transaction", "error", txErr)
			err = txErr
		}
	}()

	user, err := s.userRepository.CreateUser(ctx, tx, repository.CreateUserRequestBody(userInfo))
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return User{}, apperrors.ErrUserCreationFailed
	}

	return User(user), nil
}

func (s *service) GetLoggedInUser(ctx context.Context) (User, error) {
	tx, err := s.userRepository.BeginTx(ctx)
	if err != nil {
		slog.Error("failed to start user creation", "error", err)
		return User{}, err
	}

	defer func() {
		if txErr := s.userRepository.HandleTransaction(ctx, tx, err); txErr != nil {
			slog.Error("failed to handle transaction", "error", txErr)
			err = txErr
		}
	}()

	userIdValue := ctx.Value(constants.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return User{}, err
	}
	user, err := s.userRepository.GetUserByUserId(ctx, nil, userId)
	if err != nil {
		slog.Error("failed to get logged in user", "error", err)
		return User{}, err
	}

	return User(user), nil
}

func (s *service) UpdateUserEmail(ctx context.Context, email string) error {
	tx, err := s.userRepository.BeginTx(ctx)
	if err != nil {
		slog.Error("failed to start user creation", "error", err)
		return err
	}

	defer func() {
		if txErr := s.userRepository.HandleTransaction(ctx, tx, err); txErr != nil {
			slog.Error("failed to handle transaction", "error", txErr)
			err = txErr
		}
	}()

	userIdValue := ctx.Value(constants.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return err
	}

	err = s.userRepository.UpdateUserEmail(ctx, nil, userId, email)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		return err
	}
	return nil
}
