package auth

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	authService Service
	appConfig config.AppConfig
}

type Handler interface {
	GithubOAuthLoginUrl(w http.ResponseWriter, r *http.Request)
	GithubOAuthLoginCallback(w http.ResponseWriter, r *http.Request)
	GetLoggedInUser(w http.ResponseWriter, r *http.Request)
}

func NewHandler(authService Service, appConfig config.AppConfig) Handler {
	return &handler{
		authService: authService,
		appConfig: appConfig,
	}
}

func (h *handler) GithubOAuthLoginUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	url := h.authService.GithubOAuthLoginUrl(ctx)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *handler) GithubOAuthLoginCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := r.URL.Query().Get("code")

	token, err := h.authService.GithubOAuthLoginCallback(ctx, code)
	if err != nil {
		slog.Error("failed to login with github", "error", err)
		http.Redirect(w, r, fmt.Sprintf("%s?authError=%s", h.appConfig.ClientURL, LoginWithGithubFailed), http.StatusTemporaryRedirect)
		return
	}

	cookie := &http.Cookie{
		Name:  AccessTokenCookieName,
		Value: token,
		//TODO set domain before deploying to production
		// Domain: "yourdomain.com",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, h.appConfig.ClientURL, http.StatusPermanentRedirect)
}

func (h *handler) GetLoggedInUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userInfo, err := h.authService.GetLoggedInUser(ctx)
	if err != nil {
		slog.Error("error getting logged in user")
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "logged in user fetched successfully", userInfo)
}
