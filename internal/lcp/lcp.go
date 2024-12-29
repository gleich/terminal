package lcp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gleich/lumber/v3"
)

type LcpResponse[T any] struct {
	Data    T
	Updated time.Time
}

func fetchCache[T any](name string) (LcpResponse[T], error) {
	var zeroValue LcpResponse[T] // acts a "nil" value to be returned when there is an error
	url, err := url.JoinPath("https://lcp.dev.mattglei.ch", name, "/cache")
	if err != nil {
		return zeroValue, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		lumber.Error(err, "creating request failed")
		return zeroValue, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		lumber.Error(err, "sending request failed")
		return zeroValue, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lumber.Error(err, "reading response body failed")
		return zeroValue, err
	}

	var response LcpResponse[T]
	err = json.Unmarshal(body, &response)
	if err != nil {
		lumber.Error(err, "failed to parse json")
		lumber.Debug(string(body))
		return zeroValue, err
	}
	return response, nil
}
