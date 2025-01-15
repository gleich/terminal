package lcp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"pkg.mattglei.ch/timber"
)

type response[T any] struct {
	Data    T
	Updated time.Time
}

func fetchCache[T any](name string) (response[T], error) {
	var zeroValue response[T] // acts a "nil" value to be returned when there is an error
	url, err := url.JoinPath("https://lcp.dev.mattglei.ch", name)
	if err != nil {
		return zeroValue, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		timber.Error(err, "creating request failed")
		return zeroValue, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("LCP_TOKEN")))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		timber.Error(err, "sending request failed")
		return zeroValue, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		timber.Error(err, "reading response body failed")
		return zeroValue, err
	}

	var response response[T]
	err = json.Unmarshal(body, &response)
	if err != nil {
		timber.Error(err, "failed to parse json")
		timber.Debug(string(body))
		return zeroValue, err
	}
	return response, nil
}
