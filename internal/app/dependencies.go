package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
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
	RepositoryHandler   repoService.Handler
	AppCfg              config.AppConfig
	Client              config.Bigquery
}

func InitDependencies(db *sqlx.DB, appCfg config.AppConfig) Dependencies {
	userRepository := repository.NewUserRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService(userService, appCfg)

	authHandler := auth.NewHandler(authService, appCfg)
	userHandler := user.NewHandler(userService)
	repositoryHandler := repoService.NewHandler(repositoryService)
	contributionHandler := contribution.NewHandler(contributionService)

	return Dependencies{
		AuthService:         authService,
		UserService:         userService,
		AuthHandler:         authHandler,
		UserHandler:         userHandler,
		RepositoryHandler:   repositoryHandler,
		ContributionHandler: contributionHandler,
		AppCfg:              appCfg,
		Client:              client,
	}
}
