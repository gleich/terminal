package lcp

import (
	"pkg.mattglei.ch/lcp-2/pkg/models"
)

func FetchActivities() (response[[]models.StravaActivity], error) {
	var zeroValue response[[]models.StravaActivity]
	activities, err := fetchCache[[]models.StravaActivity]("strava")
	if err != nil {
		return zeroValue, err
	}
	return activities, nil
}
