package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"os/signal"
	"syscall"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/app"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/cronJob"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/db"
	"github.com/rs/cors"
	"github.com/urfave/cli"
)

func main() {
	cfg, err := config.LoadAppConfig()
	if err != nil {
		slog.Error("error loading app config", "error", err)
		return
	}

	cliApp := cli.NewApp()
	cliApp.Name = cfg.AppName
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start HTTP server",
			Action: func(c *cli.Context) error {
				return startApp(cfg)
			},
		},
		{
			Name:  "migrate",
			Usage: "Database migrations",
			Subcommands: []cli.Command{
				{
					Name:  "up",
					Usage: "Apply migrations",
					Action: func(c *cli.Context) error {
						m, _ := db.InitMainDBMigrations(cfg)
						m.MigrationsUp(c.Args().First())
						return nil
					},
				},
				{
					Name:  "down",
					Usage: "Rollback migrations",
					Action: func(c *cli.Context) error {
						m, _ := db.InitMainDBMigrations(cfg)
						m.MigrationsDown(c.Args().First())
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "Create a new migration file",
					Action: func(c *cli.Context) error {
						m, _ := db.InitMainDBMigrations(cfg)
						return m.CreateMigrationFile(c.Args().First())
					},
				},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func startApp(cfg config.AppConfig) error {
	ctx := context.Background()

	slog.Info("Starting CodeCuriosity Application...")

	db, err := config.InitDataStore(cfg)
	if err != nil {
		slog.Error("error initializing database", "error", err)
		return err
	}
	defer db.Close()

	bigqueryInstance, err := config.BigqueryInit(ctx, cfg)
	if err != nil {
		slog.Error("error initializing bigquery", "error", err)
		return err
	}

	httpClient := &http.Client{}

	dependencies := app.InitDependencies(db, cfg, bigqueryInstance, httpClient)

	router := app.NewRouter(dependencies)

	newCronSchedular := cronJob.NewCronSchedular()
	newCronSchedular.InitCronJobs(dependencies.ContributionService, dependencies.UserService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPServer.Port),
		Handler: router,
	}

	server.Handler = c.Handler(server.Handler)

	serverRunning := make(chan os.Signal, 1)

	signal.Notify(
		serverRunning,
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGBUS,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		slog.Info("server listening at", "port", cfg.HTTPServer.Port)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			serverRunning <- syscall.SIGINT
		}
	}()

	<-serverRunning

	slog.Info("shutting down the server")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("cannot shut HTTP server down gracefully", "error", err)
	}

	slog.Info("server shutdown successfully")
	return nil
}