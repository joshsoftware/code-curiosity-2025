package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/auth"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/badge"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
)

type Dependencies struct {
	AuthService auth.Service
	UserService user.Service
	AuthHandler auth.Handler
	UserHandler user.Handler
	BadgeHandler badge.BadgeHandler
	AppCfg      config.AppConfig
}

func InitDependencies(db *sqlx.DB, appCfg config.AppConfig) Dependencies {
	userRepository := repository.NewUserRepository(db)
	badgeRepository := repository.NewBadgeRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService(userService, appCfg)
	badgeService := badge.NewBadgeService(badgeRepository)

	authHandler := auth.NewHandler(authService, appCfg)
	userHandler := user.NewHandler(userService)
	badgeHandler := badge.NewBadgeHandler(badgeService)

	return Dependencies{
		AuthService: authService,
		UserService: userService,
		AuthHandler: authHandler,
		UserHandler: userHandler,
		BadgeHandler: badgeHandler,
		AppCfg:      appCfg,
	}
}
