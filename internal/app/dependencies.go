package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/goal"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type Dependencies struct {
	AuthService auth.Service
	UserService user.Service
	AuthHandler auth.Handler
	UserHandler user.Handler
	AppCfg      config.AppConfig
}

func InitDependencies(db *sqlx.DB, appCfg config.AppConfig) Dependencies {
	userRepository := repository.NewUserRepository(db)
	goalRepository := repository.NewGoalRepository(db)

	goalService := goal.NewGoalService(goalRepository)
	userService := user.NewService(userRepository, goalService)
	authService := auth.NewService(userService, appCfg)

	authHandler := auth.NewHandler(authService, appCfg)
	userHandler := user.NewHandler(userService)

	return Dependencies{
		AuthService: authService,
		UserService: userService,
		AuthHandler: authHandler,
		UserHandler: userHandler,
		AppCfg:      appCfg,
	}
}
