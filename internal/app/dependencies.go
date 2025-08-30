package app

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/bigquery"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/contribution"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/github"
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
	AppCfg              config.AppConfig
	Client              config.Bigquery
}

func InitDependencies(db *sqlx.DB, appCfg config.AppConfig, client config.Bigquery, httpClient *http.Client) Dependencies {
	userRepository := repository.NewUserRepository(db)
	contributionRepository := repository.NewContributionRepository(db)
	repositoryRepository := repository.NewRepositoryRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	userService := user.NewService(userRepository)
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

	return Dependencies{
		ContributionService: contributionService,
		UserService:         userService,
		AuthHandler:         authHandler,
		UserHandler:         userHandler,
		RepositoryHandler:   repositoryHandler,
		ContributionHandler: contributionHandler,
		AppCfg:              appCfg,
		Client:              client,
	}
}
