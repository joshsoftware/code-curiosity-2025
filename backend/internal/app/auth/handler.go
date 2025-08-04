package auth

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

type handler struct {
	authService Service
	appConfig   config.AppConfig
}

type Handler interface {
	GithubOAuthLoginUrl(w http.ResponseWriter, r *http.Request)
	GithubOAuthLoginCallback(w http.ResponseWriter, r *http.Request)
	GetLoggedInUser(w http.ResponseWriter, r *http.Request)
	LoginAdmin(w http.ResponseWriter, r *http.Request)
}

func NewHandler(authService Service, appConfig config.AppConfig) Handler {
	return &handler{
		authService: authService,
		appConfig:   appConfig,
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

	userIdValue := ctx.Value(middleware.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		slog.Error("error obtaining user id from context")
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	userInfo, err := h.authService.GetLoggedInUser(ctx, userId)
	if err != nil {
		slog.Error("error getting logged in user")
		status, errorMessage := apperrors.MapError(err)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "logged in user fetched successfully", userInfo)
}

func (h *handler) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var requestBody AdminLoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		slog.Error(apperrors.ErrFailedMarshal.Error(), "error", err)
		response.WriteJson(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error(), nil)
		return
	}

	adminInfo, err := h.authService.VerifyAdminCredentials(ctx, requestBody)
	if err != nil {
		slog.Error("failed to verify admin credentials", "error", err)
		status, errorMessage := apperrors.MapError(apperrors.ErrContextValue)
		response.WriteJson(w, status, errorMessage, nil)
		return
	}

	response.WriteJson(w, http.StatusOK, "admin logged in successfully", adminInfo)
}
