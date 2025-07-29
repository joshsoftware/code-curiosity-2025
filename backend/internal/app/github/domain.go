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
	LanguagesURL    string    `json:"languagesUrl"`
	UpdateDate      time.Time `json:"updatedAt"`
	RepoOwnerName   RepoOwner `json:"owner"`
	ContributorsUrl string    `json:"contributorsUrl"`
	RepoUrl         string    `json:"repoUrl"`
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
	AvatarUrl     string `json:"avatarUrl"`
	GithubUrl     string `json:"githubUrl"`
	Contributions int    `json:"contributions"`
}
