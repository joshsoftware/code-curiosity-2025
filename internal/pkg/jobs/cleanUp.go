package jobs

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/repository"
	"github.com/robfig/cron/v3"
)

func PermanentDeleteJob(db *sqlx.DB) {
	slog.Info("entering into the cleanup job")
	c := cron.New()
	_, err := c.AddFunc("36 00 * * *", func() {
		slog.Info("Job scheduled for user cleanup from database")
		ur := repository.NewUserRepository(db) // pass in *sql.DB or whatever is needed
		err := ur.DeleteUser(nil)
		if err != nil {
			slog.Error("Cleanup job error", "error", err)
		} else {
			slog.Info("User cleanup Job completed.")
		}
	})

	if err != nil {
		slog.Error("failed to start user delete job ", "error", err)
	}

	c.Start()
}
