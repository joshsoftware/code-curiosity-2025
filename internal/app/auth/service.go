package auth

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type service struct {
	githubOAuth2 oauth2.Config
	userService  user.Service
}

type Service interface {
	GithubOAuthLoginUrl(ctx context.Context) string
	GithubOAuthLoginCallback(ctx context.Context, code string) (string, error)
}

func NewService(userService user.Service) Service {
	appCfg := config.GetAppConfig()

	oauth2Config := oauth2.Config{
		ClientID:     appCfg.GithubOauth.ClientID,
		ClientSecret: appCfg.GithubOauth.ClientSecret,
		RedirectURL:  appCfg.GithubOauth.RedirectURL,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"read:user", "user:email"},
	}

	return &service{
		githubOAuth2: oauth2Config,
		userService:  userService,
	}
}

func (s *service) GithubOAuthLoginUrl(ctx context.Context) string {
	return s.githubOAuth2.AuthCodeURL("state", oauth2.AccessTypeOffline)
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

	var userInfo User
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		slog.Error("failed to unmarshal user info", "error", err)
		return "", apperrors.ErrInternalServer
	}

	_, err = s.userService.GetUserByGithubId(ctx, userInfo.GithubId)
	if err != nil {
		_, err = s.userService.CreateUser(ctx, user.CreateUserRequestBody(userInfo))
		if err != nil {
			slog.Error("failed to create user", "error", err)
			return "", apperrors.ErrUserCreationFailed
		}
	}

	jwtToken, err := jwt.GenerateJWT(userInfo.UserId, userInfo.IsAdmin)
	if err != nil {
		slog.Error("error generating jwt", "error", err)
		return "", apperrors.ErrInternalServer
	}

	return jwtToken, nil
}
