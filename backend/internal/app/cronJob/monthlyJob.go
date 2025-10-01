package cronJob

import (
	"context"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/goal"
)

type MonthlyJob struct {
	CronJob
	goalService goal.Service
}

func NewMonthlyJob(goalService goal.Service) *MonthlyJob {
	return &MonthlyJob{
		goalService: goalService,
		CronJob:     CronJob{Name: "Update User Goal Status"},
	}
}

func (m *MonthlyJob) Schedule(s *CronSchedular) error {
	_, err := s.cron.AddFunc("0 10 2 * *", func() { m.Execute(context.Background(), m.run) })
	if err != nil {
		return err
	}

	return nil
}

func (m *MonthlyJob) run(ctx context.Context) {
	m.goalService.UpdateUserGoalStatusMonthly(ctx)
}
