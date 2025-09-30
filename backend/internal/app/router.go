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
	router.HandleFunc("GET /api/v1/auth/user", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.AuthHandler.GetLoggedInUser), deps.AppCfg))

	router.HandleFunc("PATCH /api/v1/user/email", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.UserHandler.UpdateUserEmail), deps.AppCfg))
	router.HandleFunc("DELETE /api/v1/user/delete/{user_id}", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.UserHandler.SoftDeleteUser), deps.AppCfg))

	router.HandleFunc("GET /api/v1/user/contributions/all", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.ContributionHandler.FetchUserContributions), deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/overview", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.ContributionHandler.ListMonthlyContributionSummary), deps.AppCfg))
	router.HandleFunc("GET /api/v1/contributions/types", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.ContributionHandler.ListAllContributionTypes), deps.AppCfg))

	router.HandleFunc("GET /api/v1/user/repositories", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.RepositoryHandler.FetchUsersContributedRepos), deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/repositories/{repo_id}", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.RepositoryHandler.FetchParticularRepoDetails), deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/repositories/contributions/recent/{repo_id}", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.RepositoryHandler.FetchUserContributionsInRepo), deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/repositories/languages/{repo_id}", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.RepositoryHandler.FetchLanguagePercentInRepo), deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/repositories/contributors/{repo_id}", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.RepositoryHandler.FetchParticularRepoContributors), deps.AppCfg))

	router.HandleFunc("GET /api/v1/leaderboard", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.UserHandler.ListUserRanks), deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/leaderboard", middleware.Authentication(middleware.AuthorizeUnblockedUser(deps.UserHandler.GetCurrentUserRank), deps.AppCfg))

	router.HandleFunc("GET /api/v1/goal/level", middleware.Authentication(deps.GoalHandler.ListGoalLevels, deps.AppCfg))
	router.HandleFunc("POST /api/v1/goal/level/targets", middleware.Authentication(deps.GoalHandler.FetchGoalLevelTargetByGoalLevel, deps.AppCfg))
	router.HandleFunc("POST /api/v1/user/goal/level", middleware.Authentication(deps.GoalHandler.CreateUserGoalInProgress, deps.AppCfg))
	router.HandleFunc("POST /api/v1/user/goal/level/reset", middleware.Authentication(deps.GoalHandler.ResetUserCurrentGoalStatus, deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/goal/level", middleware.Authentication(deps.GoalHandler.GetUserCurrentGoalStatus, deps.AppCfg))
	router.HandleFunc("GET /api/v1/user/goal/summary", middleware.Authentication(deps.GoalHandler.FetchUserMonthlyGoalSummary, deps.AppCfg))

	router.HandleFunc("GET /api/v1/user/badges", middleware.Authentication(deps.BadgeHandler.GetBadgeDetailsOfUser, deps.AppCfg))

	router.HandleFunc("POST /api/v1/auth/admin", deps.AuthHandler.LoginAdmin)
	router.HandleFunc("PATCH /api/v1/contributions/scores/configure", middleware.Authentication(middleware.AuthorizeAdmin(deps.ContributionHandler.ConfigureContributionTypeScore), deps.AppCfg))
	router.HandleFunc("GET /api/v1/users", middleware.Authentication(middleware.AuthorizeAdmin(deps.UserHandler.ListAllUsers), deps.AppCfg))
	router.HandleFunc("PATCH /api/v1/users/{user_id}", middleware.Authentication(middleware.AuthorizeAdmin(deps.UserHandler.BlockOrUnblockUser), deps.AppCfg))

	return middleware.CorsMiddleware(router, deps.AppCfg)
}
