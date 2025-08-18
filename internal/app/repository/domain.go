package repository

import "time"

type RepoOWner struct {
	Login string `json:"login"`
}

type FetchRepositoryDetailsResponse struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	LanguagesURL    string    `json:"languages_url"`
	UpdateDate      time.Time `json:"updated_at"`
	RepoOwnerName   RepoOWner `json:"owner"`
	ContributorsUrl string    `json:"contributors_url"`
	RepoUrl         string    `json:"html_url"`
}

type Repository struct {
	Id              int
	GithubRepoId    int
	RepoName        string
	Description     string
	LanguagesUrl    string
	RepoUrl         string
	OwnerName       string
	UpdateDate      time.Time
	ContributorsUrl string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
