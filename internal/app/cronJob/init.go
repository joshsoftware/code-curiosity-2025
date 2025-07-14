package cronJob

import (
	"log/slog"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/contribution"
	"github.com/robfig/cron/v3"
)

type CronSchedular struct {
	cron *cron.Cron
}

func NewCronSchedular() *CronSchedular {
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		slog.Error("failed to load IST timezone", "error", err)
	}

	return &CronSchedular{
		cron: cron.New(cron.WithLocation(location)),
	}
}

func (c *CronSchedular) InitCronJobs(contributionService contribution.Service) {
	jobs := []Job{
		NewDailyJob(contributionService),
	}

	for _, job := range jobs {
		if err := job.Schedule(c); err != nil {
			slog.Error("failed to execute cron job")
		}
	}

	c.cron.Start()
}
