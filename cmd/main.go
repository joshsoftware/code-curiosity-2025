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
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
)

func main() {
	ctx := context.Background()

	cfg,err := config.LoadAppConfig()
	if err != nil {
		slog.Error("error loading app config", "error", err)
		return
	}


	db, err := config.InitDataStore(cfg)
	if err != nil {
		slog.Error("error initializing database", "error", err)
		return
	}
	defer db.Close()

	dependencies := app.InitDependencies(db,cfg)

	router := app.NewRouter(dependencies)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPServer.Port),
		Handler: router,
	}

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
}
