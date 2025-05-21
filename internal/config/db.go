package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDataStore(appCfg AppConfig) (*sqlx.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appCfg.Database.Host, appCfg.Database.Port, appCfg.Database.User, appCfg.Database.Password, appCfg.Database.Name)

	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
