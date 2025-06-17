package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/bigquery"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/contribution"
	repoService "github.com/joshsoftware/code-curiosity-2025/internal/app/repository"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"

	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type Dependencies struct {
	AuthService         auth.Service
	UserService         user.Service
	AuthHandler         auth.Handler
	UserHandler         user.Handler
	ContributionHandler contribution.Handler
	AppCfg              config.AppConfig
	Client              config.Bigquery
}

func InitDependencies(db *sqlx.DB, appCfg config.AppConfig, client config.Bigquery) Dependencies {
	userRepository := repository.NewUserRepository(db)
	contributionRepository := repository.NewContributionRepository(db)
	repositoryRepository := repository.NewRepositoryRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService(userService, appCfg)
	bigqueryService := bigquery.NewService(client, userRepository)
	repositoryService := repoService.NewService(repositoryRepository, appCfg)
	contributionService := contribution.NewService(bigqueryService, contributionRepository, repositoryService, userService)

	authHandler := auth.NewHandler(authService, appCfg)
	userHandler := user.NewHandler(userService)
	contributionHandler := contribution.NewHandler(contributionService)

	return Dependencies{
		AuthService:         authService,
		UserService:         userService,
		AuthHandler:         authHandler,
		UserHandler:         userHandler,
		ContributionHandler: contributionHandler,
		AppCfg:              appCfg,
		Client:              client,
	}
}
