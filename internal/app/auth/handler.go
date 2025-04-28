package auth

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/config"
)

func GithubOAuthLoginUrl(authService Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		url := authService.GithubOAuthLoginUrl(ctx)

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func GithubOAuthLoginCallback(authService Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		appCfg := config.GetAppConfig()
		code := r.URL.Query().Get("code")

		token, err := authService.GithubOAuthLoginCallback(ctx, code)
		if err != nil {
			slog.Error("failed to login with github", "error", err)
			http.Redirect(w, r, fmt.Sprintf("%s?authError=%s", appCfg.ClientURL, LoginWithGithubFailed), http.StatusTemporaryRedirect)
			return
		}

		cookie:= &http.Cookie{
			Name: AccessTokenCookieName,
			Value: token,
			//TODO set domain before deploying to production
			// Domain: "yourdomain.com",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, appCfg.ClientURL, http.StatusPermanentRedirect)
	}
}
