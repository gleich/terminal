package lcp

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gleich/lumber/v3"
)

type SteamGame struct {
	Name                string              `json:"name"`
	AppID               int32               `json:"app_id"`
	IconURL             string              `json:"icon_url"`
	RTimeLastPlayed     time.Time           `json:"rtime_last_played"`
	PlaytimeForever     int32               `json:"playtime_forever"`
	URL                 string              `json:"url"`
	HeaderURL           string              `json:"header_url"`
	LibraryURL          *string             `json:"library_url"`
	LibraryHeroURL      string              `json:"library_hero_url"`
	LibraryHeroLogoURL  string              `json:"library_hero_logo_url"`
	AchievementProgress *float32            `json:"achievement_progress"`
	Achievements        *[]SteamAchievement `json:"achievements"`
}

type SteamAchievement struct {
	ApiName     string     `json:"api_name"`
	Achieved    bool       `json:"achieved"`
	Icon        string     `json:"icon"`
	DisplayName string     `json:"display_name"`
	Description *string    `json:"description"`
	UnlockTime  *time.Time `json:"unlock_time"`
}

func FetchGames() (Response[[]SteamGame], error) {
	req, err := http.NewRequest("GET", "https://lcp.dev.mattglei.ch/steam/cache", nil)
	if err != nil {
		lumber.Error(err, "creating new request failed")
		return Response[[]SteamGame]{}, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("LCP_ACCESS_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		lumber.Error(err, "sending request for steam games failed")
		return Response[[]SteamGame]{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		lumber.Error(err, "reading response body failed")
		return Response[[]SteamGame]{}, err
	}

	var response Response[[]SteamGame]
	err = json.Unmarshal(body, &response)
	if err != nil {
		lumber.Error(err, "failed to parse json")
		lumber.Debug(string(body))
		return Response[[]SteamGame]{}, err
	}
	return response, nil
}
