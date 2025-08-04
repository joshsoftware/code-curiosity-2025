package user

import (
	"context"
	"log/slog"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/goal"
	repoService "github.com/joshsoftware/code-curiosity-2025/internal/app/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type service struct {
	userRepository    repository.UserRepository
	goalService       goal.Service
	repositoryService repoService.Service
}

type Service interface {
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByGithubId(ctx context.Context, githubId int) (User, error)
	CreateUser(ctx context.Context, userInfo CreateUserRequestBody) (User, error)
	UpdateUserEmail(ctx context.Context, userId int, email string) error
	SoftDeleteUser(ctx context.Context, userId int) error
	HardDeleteUsers(ctx context.Context) error
	RecoverAccountInGracePeriod(ctx context.Context, userID int) error
	UpdateUserCurrentBalance(ctx context.Context, transaction Transaction) error
	GetAllUsersRank(ctx context.Context) ([]LeaderboardUser, error)
	GetCurrentUserRank(ctx context.Context, userId int) (LeaderboardUser, error)
	UpdateCurrentActiveGoalId(ctx context.Context, userId int, level string) (int, error)
	GetLoggedInAdmin(ctx context.Context, adminInfo AdminLoginRequest) (User, error)
	ListAllUsers(ctx context.Context) ([]User, error)
	BlockOrUnblockUser(ctx context.Context, userID int, block bool) error
}

func NewService(userRepository repository.UserRepository, goalService goal.Service, repositoryService repoService.Service) Service {
	return &service{
		userRepository:    userRepository,
		goalService:       goalService,
		repositoryService: repositoryService,
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

func (s *service) UpdateUserEmail(ctx context.Context, userId int, email string) error {
	err := s.userRepository.UpdateUserEmail(ctx, nil, userId, email)
	if err != nil {
		slog.Error("failed to update user email", "error", err)
		return err
	}

	return nil
}

func (s *service) SoftDeleteUser(ctx context.Context, userID int) error {
	now := time.Now()
	err := s.userRepository.MarkUserAsDeleted(ctx, nil, userID, now)
	if err != nil {
		slog.Error("unable to softdelete user", "error", err)
		return apperrors.ErrInternalServer
	}
	return nil
}

func (s *service) HardDeleteUsers(ctx context.Context) error {
	err := s.userRepository.HardDeleteUsers(ctx, nil)
	if err != nil {
		slog.Error("error deleting users that are soft deleted for more than three months", "error", err)
		return err
	}

	return nil
}

func (s *service) RecoverAccountInGracePeriod(ctx context.Context, userID int) error {
	err := s.userRepository.RecoverAccountInGracePeriod(ctx, nil, userID)
	if err != nil {
		slog.Error("failed to recover account in grace period", "error", err)
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

func (s *service) GetAllUsersRank(ctx context.Context) ([]LeaderboardUser, error) {
	userRanks, err := s.userRepository.GetAllUsersRank(ctx, nil)
	if err != nil {
		slog.Error("error obtaining all users rank", "error", err)
		return nil, err
	}

	Leaderboard := make([]LeaderboardUser, len(userRanks))
	for i, l := range userRanks {
		userContributedReposCount, err := s.repositoryService.FetchUserContributedReposCount(ctx, l.Id)
		if err != nil {
			slog.Error("error fetching user contributed repos count", "error", err)
			return nil, err
		}

		Leaderboard[i].Id = l.Id
		Leaderboard[i].GithubUsername = l.GithubUsername
		Leaderboard[i].ContributedReposCount = userContributedReposCount
		Leaderboard[i].AvatarUrl = l.AvatarUrl
		Leaderboard[i].Rank = l.Rank
		Leaderboard[i].CurrentBalance = l.CurrentBalance
	}

	return Leaderboard, nil
}

func (s *service) GetCurrentUserRank(ctx context.Context, userId int) (LeaderboardUser, error) {
	currentUserRank, err := s.userRepository.GetCurrentUserRank(ctx, nil, userId)
	if err != nil {
		slog.Error("error obtaining current user rank", "error", err)
		return LeaderboardUser{}, err
	}

	currentUserContributedReposCount, err := s.repositoryService.FetchUserContributedReposCount(ctx, userId)
	if err != nil {
		slog.Error("error fetching user contributed repos count", "error", err)
		return LeaderboardUser{}, err
	}

	leaderboardUser := LeaderboardUser{
		Id:                    currentUserRank.Id,
		GithubUsername:        currentUserRank.GithubUsername,
		AvatarUrl:             currentUserRank.AvatarUrl,
		ContributedReposCount: currentUserContributedReposCount,
		CurrentBalance:        currentUserRank.CurrentBalance,
		Rank:                  currentUserRank.Rank,
	}
	return leaderboardUser, nil
}

func (s *service) UpdateCurrentActiveGoalId(ctx context.Context, userId int, level string) (int, error) {

	goalId, err := s.goalService.GetGoalIdByGoalLevel(ctx, level)

	if err != nil {
		slog.Error("error occured while fetching goal id by goal level")
		return 0, err
	}

	goalId, err = s.userRepository.UpdateCurrentActiveGoalId(ctx, nil, userId, goalId)

	if err != nil {
		slog.Error("failed to update current active goal id", "error", err)
	}

	return goalId, err
}

func (s *service) GetLoggedInAdmin(ctx context.Context, adminInfo AdminLoginRequest) (User, error) {
	admin, err := s.userRepository.GetAdminByCredentials(ctx, nil, repository.AdminLoginRequest(adminInfo))
	if err != nil {
		slog.Error("failed to verify admin credentials", "error", err)
		return User{}, err
	}

	return User(admin), nil
}

func (s *service) ListAllUsers(ctx context.Context) ([]User, error) {
	users, err := s.userRepository.GetAllUsers(ctx, nil)
	if err != nil {
		slog.Error("failed to fetch all users", "error", err)
		return nil, apperrors.ErrInternalServer
	}

	serviceUsers := make([]User, len(users))

	for i, u := range users {
		serviceUsers[i] = User(u)
	}

	return serviceUsers, nil
}

func (s *service) BlockOrUnblockUser(ctx context.Context, userID int, block bool) error {
	err := s.userRepository.UpdateUserBlockStatus(ctx, nil, userID, block)
	if err != nil {
		slog.Error("failed to block/unblock user", "error", err)
		return err
	}

	return nil
}
