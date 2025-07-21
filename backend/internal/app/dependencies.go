package app

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/badge"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/bigquery"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/contribution"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/github"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/goal"
	repoService "github.com/joshsoftware/code-curiosity-2025/internal/app/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/transaction"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type Dependencies struct {
	ContributionService contribution.Service
	UserService         user.Service
	AuthHandler         auth.Handler
	UserHandler         user.Handler
	ContributionHandler contribution.Handler
	RepositoryHandler   repoService.Handler
	GoalHandler         goal.Handler
	BadgeHandler        badge.Handler
	AppCfg              config.AppConfig
	Client              config.Bigquery
}

func InitDependencies(db *sqlx.DB, appCfg config.AppConfig, client config.Bigquery, httpClient *http.Client) Dependencies {
	badgeRepository := repository.NewBadgeRepository(db)
	goalRepository := repository.NewGoalRepository(db)
	userRepository := repository.NewUserRepository(db)
	contributionRepository := repository.NewContributionRepository(db)
	repositoryRepository := repository.NewRepositoryRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	badgeService := badge.NewService(badgeRepository)
	goalService := goal.NewService(goalRepository, contributionRepository, badgeService)
	userService := user.NewService(userRepository, goalService)
	authService := auth.NewService(userService, appCfg)
	bigqueryService := bigquery.NewService(client, userRepository)
	githubService := github.NewService(appCfg, httpClient)
	repositoryService := repoService.NewService(repositoryRepository, githubService)
	transactionService := transaction.NewService(transactionRepository, userService)
	contributionService := contribution.NewService(bigqueryService, contributionRepository, repositoryService, userService, transactionService, httpClient)

	authHandler := auth.NewHandler(authService, appCfg)
	userHandler := user.NewHandler(userService)
	repositoryHandler := repoService.NewHandler(repositoryService, githubService)
	contributionHandler := contribution.NewHandler(contributionService)
	goalHandler := goal.NewHandler(goalService)
	badgeHandler := badge.NewHandler(badgeService)

	return Dependencies{
		ContributionService: contributionService,
		UserService:         userService,
		AuthHandler:         authHandler,
		UserHandler:         userHandler,
		RepositoryHandler:   repositoryHandler,
		ContributionHandler: contributionHandler,
		GoalHandler:         goalHandler,
		BadgeHandler:        badgeHandler,
		AppCfg:              appCfg,
		Client:              client,
	}
}
