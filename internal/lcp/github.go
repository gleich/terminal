package lcp

import (
	"pkg.mattglei.ch/lcp-2/pkg/models"
)

func FetchRepositories() (response[[]models.GitHubRepository], error) {
	var zeroValue response[[]models.GitHubRepository]
	repos, err := fetchCache[[]models.GitHubRepository]("github")
	if err != nil {
		return zeroValue, nil
	}
	return repos, nil
}
