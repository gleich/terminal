package lcp

import (
	"time"
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

func FetchActivities() (response[[]StravaActivity], error) {
	var zeroValue response[[]StravaActivity]
	activities, err := fetchCache[[]StravaActivity]("strava")
	if err != nil {
		return zeroValue, err
	}
	return activities, nil
}
