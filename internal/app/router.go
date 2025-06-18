package app

import (
	"net/http"

	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/middleware"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/response"
)

func NewRouter(deps Dependencies) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, http.StatusOK, "Server is up and running..", nil)
	})

	router.HandleFunc("GET /api/v1/auth/github", deps.AuthHandler.GithubOAuthLoginUrl)
	router.HandleFunc("GET /api/v1/auth/github/callback", deps.AuthHandler.GithubOAuthLoginCallback)
	router.HandleFunc("GET /api/v1/auth/user", middleware.Authentication(deps.AuthHandler.GetLoggedInUser, deps.AppCfg))

	router.HandleFunc("PATCH /api/v1/user/email", middleware.Authentication(deps.UserHandler.UpdateUserEmail, deps.AppCfg))

	router.HandleFunc("GET /api/v1/user/contributions/latest", middleware.Authentication(deps.ContributionHandler.FetchUserLatestContributions, deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/contributions/recent", middleware.Authentication(deps.ContributionHandler.FetchUsersFiveRecentContributions, deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/contributions/all", middleware.Authentication(deps.ContributionHandler.FetchUsersAllContributions, deps.AppCfg))

	router.HandleFunc("GET /api/v1/user/repositories", middleware.Authentication(deps.RepositoryHandler.FetchUsersContributedRepos, deps.AppCfg))
	return middleware.CorsMiddleware(router, deps.AppCfg)
}
