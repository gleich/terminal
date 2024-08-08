package lcp

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gleich/lumber/v2"
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

func FetchRepositories() (Response[[]GitHubRepository], error) {
	req, err := http.NewRequest("GET", "https://lcp.dev.mattglei.ch/github/cache", nil)
	if err != nil {
		lumber.Error(err, "Failed to create new request")
		return Response[[]GitHubRepository]{}, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LCP_ACCESS_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		lumber.Error(err, "Failed to send request for GitHub projects")
		return Response[[]GitHubRepository]{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lumber.Error(err, "reading response body failed")
		return Response[[]GitHubRepository]{}, err
	}

	var response Response[[]GitHubRepository]
	err = json.Unmarshal(body, &response)
	if err != nil {
		lumber.Error(err, "failed to parse json")
		lumber.Debug(string(body))
		return Response[[]GitHubRepository]{}, err
	}
	return response, nil
}
