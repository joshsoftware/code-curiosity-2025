package bigquery

import "time"

type ContributionResponse struct {
	ID         string    `bigquery:"id"`
	Type       string    `bigquery:"type"`
	ActorID    int       `bigquery:"actor_id"`
	ActorLogin string    `bigquery:"actor_login"`
	RepoID     int       `bigquery:"repo_id"`
	RepoName   string    `bigquery:"repo_name"`
	RepoUrl    string    `bigquery:"repo_url"`
	Payload    string    `bigquery:"payload"`
	CreatedAt  time.Time `bigquery:"created_at"`
}
