package lcp

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gleich/lumber/v2"
)

type StravaActivity struct {
	Name               string    `json:"name"`
	SportType          string    `json:"sport_type"`
	StartDate          time.Time `json:"start_date"`
	Timezone           string    `json:"timezone"`
	MapBlurImage       *string   `json:"map_blur_image"`
	MapImageURL        *string   `json:"map_image_url"`
	HasMap             bool      `json:"has_map"`
	TotalElevationGain float32   `json:"total_elevation_gain"`
	MovingTime         uint32    `json:"moving_time"`
	Distance           float32   `json:"distance"`
	ID                 uint64    `json:"id"`
	AverageHeartrate   float32   `json:"average_heartrate"`
}

func FetchActivities() (Response[[]StravaActivity], error) {
	req, err := http.NewRequest("GET", "https://lcp.dev.mattglei.ch/strava/cache", nil)
	if err != nil {
		lumber.Error(err, "Failed to create new request")
		return Response[[]StravaActivity]{}, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LCP_ACCESS_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		lumber.Error(err, "Failed to send request for Strava activities")
		return Response[[]StravaActivity]{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lumber.Error(err, "reading response body failed")
		return Response[[]StravaActivity]{}, err
	}

	var response Response[[]StravaActivity]
	err = json.Unmarshal(body, &response)
	if err != nil {
		lumber.Error(err, "failed to parse json")
		lumber.Debug(string(body))
		return Response[[]StravaActivity]{}, err
	}
	return response, nil
}
