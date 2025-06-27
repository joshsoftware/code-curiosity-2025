package cronJob

import (
	"context"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/contribution"
)

type DailyJob struct {
	CronJob
	contributionService contribution.Service
}

func NewDailyJob(contributionService contribution.Service) *DailyJob {
	return &DailyJob{
		contributionService: contributionService,
		CronJob:             CronJob{Name: "Fetch Contributions Daily"},
	}
}

func (d *DailyJob) Schedule(s *CronSchedular) error {
	_, err := s.cron.AddFunc("0 1 * * *", func() { d.Execute(context.Background(), d.run)() })
	if err != nil {
		return err
	}

	return nil
}

func (d *DailyJob) run(ctx context.Context) {
	d.contributionService.ProcessFetchedContributions(ctx)
}
