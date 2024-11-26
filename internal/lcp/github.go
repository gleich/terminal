package lcp

import (
	"time"
)

type GitHubRepository struct {
	Name          string    `json:"name"`
	Owner         string    `json:"owner"`
	Language      string    `json:"language"`
	LanguageColor string    `json:"language_color"`
	Description   string    `json:"description"`
	UpdatedAt     time.Time `json:"updated_at"`
	Stargazers    int32     `json:"stargazers"`
	ID            string    `json:"id"`
	URL           string    `json:"url"`
}

func FetchRepositories() (LcpResponse[[]GitHubRepository], error) {
	var zeroValue LcpResponse[[]GitHubRepository]
	repos, err := fetchCache[[]GitHubRepository]("github")
	if err != nil {
		return zeroValue, nil
	}
	return repos, nil
}
