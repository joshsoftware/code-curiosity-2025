package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	_ "github.com/lib/pq"
)

func InitDataStore(appCfg config.AppConfig) (*sqlx.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appCfg.Database.Host, appCfg.Database.Port, appCfg.Database.User, appCfg.Database.Password, appCfg.Database.Name)

	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
