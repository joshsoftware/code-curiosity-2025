package cronJob

import (
	"log/slog"
	"time"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/contribution"
	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
	"github.com/robfig/cron/v3"
)

type CronSchedular struct {
	cron *cron.Cron
}

func NewCronSchedular() *CronSchedular {
	//CHANGE AND SET TO UTC TIMEZONE
	return &CronSchedular{
		cron: cron.New(cron.WithLocation(time.UTC)),
	}
}

func (c *CronSchedular) InitCronJobs(contributionService contribution.Service, userService user.Service) {
	jobs := []Job{
		NewDailyJob(contributionService),
		NewCleanupJob(userService),
	}

	for _, job := range jobs {
		if err := job.Schedule(c); err != nil {
			slog.Error("failed to execute cron job")
		}
	}

	c.cron.Start()
}
