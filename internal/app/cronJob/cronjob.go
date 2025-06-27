package cronJob

import (
	"context"
	"log/slog"
	"time"
)

type Job interface {
	Schedule(c *CronSchedular) error
}

type CronJob struct {
	Name string
}

func (c *CronJob) Execute(ctx context.Context, fn func(context.Context)) func() {
	return func() {
		slog.Info("cron job started at", "time ", time.Now())
		defer func() {
			slog.Info("cron job completed")
		}()

		fn(ctx)
	}
}
