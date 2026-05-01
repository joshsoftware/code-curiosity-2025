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
	GetContributionScores(ctx context.Context) ([]ContributionScore, error)
	UpdateContributionScore(ctx context.Context, contributionScoreId int, score int) error
	GetLeaderboard(ctx context.Context) ([]LeaderboardEntry, error)
	UpdateUserBlocked(ctx context.Context, userId int, isBlocked bool) error
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

func (s *service) GetContributionScores(ctx context.Context) ([]ContributionScore, error) {
	repositoryContributionScores, err := s.userRepository.ListContributionScores(ctx, nil)
	if err != nil {
		slog.Error("failed to get contribution scores", "error", err)
		return nil, err
	}

	contributionScores := make([]ContributionScore, len(repositoryContributionScores))
	for index, contributionScore := range repositoryContributionScores {
		contributionScores[index] = ContributionScore(contributionScore)
	}

	return contributionScores, nil
}

func (s *service) UpdateContributionScore(ctx context.Context, contributionScoreId int, score int) error {
	err := s.userRepository.UpdateContributionScore(ctx, nil, contributionScoreId, score)
	if err != nil {
		slog.Error("failed to update contribution score", "error", err)
		return err
	}

	return nil
}

func (s *service) GetLeaderboard(ctx context.Context) ([]LeaderboardEntry, error) {
	repositoryLeaderboard, err := s.userRepository.ListLeaderboard(ctx, nil)
	if err != nil {
		slog.Error("failed to get leaderboard", "error", err)
		return nil, err
	}

	leaderboard := make([]LeaderboardEntry, len(repositoryLeaderboard))
	for index, leaderboardEntry := range repositoryLeaderboard {
		leaderboard[index] = LeaderboardEntry(leaderboardEntry)
	}

	return leaderboard, nil
}

func (s *service) UpdateUserBlocked(ctx context.Context, userId int, isBlocked bool) error {
	err := s.userRepository.UpdateUserBlocked(ctx, nil, userId, isBlocked)
	if err != nil {
		slog.Error("failed to update user blocked status", "error", err)
		return err
	}

	return nil
}
