package repository

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
	"github.com/joshsoftware/code-curiosity-2025/internal/pkg/apperrors"
)

func SeedData(appCfg config.AppConfig) error {
	dbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appCfg.Database.Host,
		appCfg.Database.Port,
		appCfg.Database.User,
		appCfg.Database.Password,
		appCfg.Database.Name,
	)

	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return err
	}

	seedQueries := []string{
		// Insert default admin user
		`INSERT INTO users VALUES (
			DEFAULT, 
			0, 
			'admin', 
			'', 
			'', 
			0, 
			false, 
			true, 
			'$2a$14$gWxgkAc0uPTxkSBlMTudZusI/4QmQQssMXW8NjjZTJqsDx7PKdBvG', 
			false, 
			DEFAULT, 
			DEFAULT, 
			DEFAULT
		)`,

		// Contribution scores
		`INSERT INTO contribution_score (admin_id, contribution_type, score)
		VALUES 
			(1, 'CommitAdded', 10),
			(1, 'IssueOpened', 10),
			(1, 'IssueClosed', 20),
			(1, 'IssueCompleted', 30),
			(1, 'PullRequestOpened', 40),
			(1, 'PullRequestMerged', 50),
			(1, 'PullRequestUpdated', 60),
			(1, 'IssueComment', 10),
			(1, 'PullRequestComment', 10)`,

		// Goal levels
		`INSERT INTO goal_level (level)
		VALUES 
			('Custom'),
			('Beginner'),
			('Intermediate'),
			('Advanced')`,

		// Goal level targets
		`INSERT INTO goal_level_target (goal_level_id, contribution_score_id, target)
		VALUES
			(2, 5, 1),

			(3, 1, 2),
			(3, 2, 3),
			(3, 3, 4),
			(3, 4, 1),
			(3, 5, 2),
			(3, 6, 3),
			(3, 7, 4),
			(3, 8, 1),
			(3, 9, 2),

			(4, 1, 3),
			(4, 2, 4),
			(4, 3, 1),
			(4, 4, 2),
			(4, 5, 3),
			(4, 6, 4),
			(4, 7, 1),
			(4, 8, 2),
			(4, 9, 3)`,
	}

	for _, query := range seedQueries {
		_, err := db.Exec(query)
		if err != nil {
			slog.Error("Failed to execute seed query",
				"query", query,
				"error", err,
			)
			return apperrors.ErrInternalServer
		}
	}

	slog.Info("Seed data loaded successfully.")
	return nil
}
