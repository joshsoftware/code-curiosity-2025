package github

import "time"

const AuthorizationKey = "Authorization"

type RepoOwner struct {
	Login string `json:"login"`
}

type FetchRepositoryDetailsResponse struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	LanguagesURL    string    `json:"languages_url"`
	UpdateDate      time.Time `json:"updated_at"`
	RepoOwnerName   RepoOwner `json:"owner"`
	ContributorsUrl string    `json:"contributors_url"`
	RepoUrl         string    `json:"html_url"`
}

type RepoLanguages map[string]int

type RepoContributorsResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"login"`
	AvatarUrl     string `json:"avatar_url"`
	GithubUrl     string `json:"html_url"`
	Contributions int    `json:"contributions"`
}

type FetchRepositoryContributorsResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	AvatarUrl     string `json:"avatar_url"`
	GithubUrl     string `json:"github_url"`
	Contributions int    `json:"contributions"`
}
