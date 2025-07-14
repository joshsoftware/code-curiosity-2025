package bigquery

import "time"

const DailyQuery = `SELECT 
  id,
  type,
  public,
  actor.id AS actor_id,
  actor.login AS actor_login,
  actor.gravatar_id AS actor_gravatar_id,
  actor.url AS actor_url,
  actor.avatar_url AS actor_avatar_url,
  repo.id AS repo_id,
  repo.name AS repo_name,
  repo.url AS repo_url,
  payload,
  created_at,
  other
FROM 
  githubarchive.day.%s
WHERE 
  type IN (
    'IssuesEvent', 
    'PullRequestEvent', 
    'PullRequestReviewEvent', 
    'IssueCommentEvent', 
    'PullRequestReviewCommentEvent'
  )
  AND (
    actor.id IN (%s) 
  )`

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
