package app

import (
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

func NewRouter(deps Dependencies) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "Server is up and running..", nil)
	})

	router.HandleFunc("GET /api/v1/auth/github", auth.GithubOAuthLoginUrl(deps.AuthService))
	router.HandleFunc("GET /api/v1/auth/github/callback", auth.GithubOAuthLoginCallback(deps.AuthService))
	router.HandleFunc("GET /api/v1/auth/user", middleware.Authentication(auth.GetLoggedInUser(deps.AuthService)))

	router.HandleFunc("PATCH /api/v1/user/email", user.UpdateUserEmail(deps.UserService))

	return middleware.CorsMiddleware(router)
}
