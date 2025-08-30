package cronJob

import (
	"context"

	"github.com/joshsoftware/code-curiosity-2025/internal/app/user"
)

type CleanupJob struct {
	CronJob
	userService user.Service
}

func NewCleanupJob(userService user.Service) *CleanupJob {
	return &CleanupJob{
		userService: userService,
		CronJob:     CronJob{Name: "User Cleanup Job Daily"},
	}
}

func (c *CleanupJob) Schedule(s *CronSchedular) error {
	_, err := s.cron.AddFunc("00 18 * * *", func() { c.Execute(context.Background(), c.run) })
	if err != nil {
		return err
	}

	return nil
}

func (c *CleanupJob) run(ctx context.Context) {
	c.userService.HardDeleteUsers(ctx)
}
