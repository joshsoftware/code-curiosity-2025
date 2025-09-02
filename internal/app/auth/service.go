package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/jwt"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type service struct {
	githubOAuth2 oauth2.Config
	userService  user.Service
	appCfg       config.AppConfig
}

type Service interface {
	GithubOAuthLoginUrl(ctx context.Context) string
	GithubOAuthLoginCallback(ctx context.Context, code string) (string, error)
	GetLoggedInUser(ctx context.Context) (User, error)
}

func NewService(userService user.Service, appCfg config.AppConfig) Service {
	oauth2Config := oauth2.Config{
		ClientID:     appCfg.GithubOauth.ClientID,
		ClientSecret: appCfg.GithubOauth.ClientSecret,
		RedirectURL:  appCfg.GithubOauth.RedirectURL,
		Endpoint:     github.Endpoint,
		Scopes:       []string{GithubOauthScope},
	}

	return &service{
		githubOAuth2: oauth2Config,
		userService:  userService,
		appCfg:       appCfg,
	}
}

func (s *service) GithubOAuthLoginUrl(ctx context.Context) string {
	return s.githubOAuth2.AuthCodeURL(GitHubOAuthState, oauth2.AccessTypeOffline)
}

func (s *service) GithubOAuthLoginCallback(ctx context.Context, code string) (string, error) {
	token, err := s.githubOAuth2.Exchange(ctx, code)
	if err != nil {
		slog.Error("failed to exchange token", "error", err)
		return "", apperrors.ErrGithubTokenExchangeFailed
	}

	client := s.githubOAuth2.Client(ctx, token)
	resp, err := client.Get(GetUserGithubUrl)
	if err != nil {
		slog.Error("failed to get user info", "error", err)
		return "", apperrors.ErrFailedToGetGithubUser
	}
	defer resp.Body.Close()

	var userInfo GithubUserResponse
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		slog.Error("failed to unmarshal user info", "error", err)
		return "", apperrors.ErrInternalServer
	}

	userData, err := s.userService.GetUserByGithubId(ctx, userInfo.GithubId)
	if err != nil {
		userData, err = s.userService.CreateUser(ctx, user.CreateUserRequestBody(userInfo))
		if err != nil {
			slog.Error("failed to create user", "error", err)
			return "", apperrors.ErrUserCreationFailed
		}
	}

	jwtToken, err := jwt.GenerateJWT(userData.Id, userInfo.IsAdmin, s.appCfg)
	if err != nil {
		slog.Error("error generating jwt", "error", err)
		return "", apperrors.ErrInternalServer
	}

	// soft delete checker
	err = s.userService.RecoverAccountInGracePeriod(ctx, userData.Id)
	if err != nil {
		slog.Error("error in recovering account in grace period during login", "error", err)
		return "", apperrors.ErrInternalServer
	}
	// token print

	fmt.Println(jwtToken)

	return jwtToken, nil
}

func (s *service) GetLoggedInUser(ctx context.Context) (User, error) {
	userIdValue := ctx.Value(middleware.UserIdKey)

	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		return User{}, apperrors.ErrInternalServer
	}

	user, err := s.userService.GetUserById(ctx, userId)
	if err != nil {
		slog.Error("failed to get logged in user", "error", err)
		return User{}, apperrors.ErrInternalServer
	}

	return User(user), nil
}
