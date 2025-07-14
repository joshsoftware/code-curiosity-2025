package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

type HTTPServer struct {
	Port string `yaml:"port" required:"true"`
}

type Database struct {
	Host     string `yaml:"host" required:"true"`
	Port     int    `yaml:"port" required:"true"`
	User     string `yaml:"user" required:"true"`
	Password string `yaml:"password" required:"true"`
	Name     string `yaml:"name" required:"true"`
}

type GithubOauth struct {
	ClientID     string `yaml:"client_id" required:"true"`
	ClientSecret string `yaml:"client_secret" required:"true"`
	RedirectURL  string `yaml:"redirect_url" required:"true"`
}

type BigqueryProject struct {
	ProjectID string `yaml:"project_id" required:"true"`
}

type AppConfig struct {
	IsProduction              bool            `yaml:"is_production"`
	HTTPServer                HTTPServer      `yaml:"http_server"`
	Database                  Database        `yaml:"database"`
	JWTSecret                 string          `yaml:"jwt_secret"`
	ClientURL                 string          `yaml:"client_url"`
	GithubOauth               GithubOauth     `yaml:"github_oauth"`
	BigqueryProject           BigqueryProject `yaml:"bigquery_project"`
	GithubPersonalAccessToken string          `yaml:"github_personal_access_token"`
}

func LoadAppConfig() (AppConfig, error) {
	appConfigPath := os.Getenv("CONFIG_PATH")

	if appConfigPath == "" {
		return AppConfig{}, apperrors.ErrNoAppConfigPath
	}

	if _, err := os.Stat(appConfigPath); os.IsNotExist(err) {
		return AppConfig{}, apperrors.ErrNoAppConfigPath
	}

	var appCfg AppConfig
	if err := cleanenv.ReadConfig(appConfigPath, &appCfg); err != nil {
		return AppConfig{}, apperrors.ErrFailedToLoadAppConfig
	}

	return appCfg, nil
}
