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
	Name      string    `json:"name"`
	SportType string    `json:"sport_type"`
	StartDate time.Time `json:"start_date"`
	Timezone  string    `json:"timezone"`
	Map       struct {
		SummaryPolyline string  `json:"summary_polyline"`
		MapBlurImage    *string `json:"map_blur_image"`
	} `json:"map"`
	Trainer            bool    `json:"trainer"`
	Commute            bool    `json:"commute"`
	Private            bool    `json:"private"`
	AverageSpeed       float32 `json:"average_speed"`
	MaxSpeed           float32 `json:"max_speed"`
	AverageTemp        int32   `json:"average_temp,omitempty"`
	AverageCadence     float32 `json:"average_cadence,omitempty"`
	AverageWatts       float32 `json:"average_watts,omitempty"`
	DeviceWatts        bool    `json:"device_watts,omitempty"`
	AverageHeartrate   float32 `json:"average_heartrate,omitempty"`
	TotalElevationGain float32 `json:"total_elevation_gain"`
	MovingTime         uint32  `json:"moving_time"`
	SufferScore        float32 `json:"suffer_score,omitempty"`
	PrCount            uint32  `json:"pr_count"`
	Distance           float32 `json:"distance"`
	ID                 uint64  `json:"id"`
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
