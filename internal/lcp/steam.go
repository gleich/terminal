package lcp

import (
	"time"
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

func FetchGames() (LcpResponse[[]SteamGame], error) {
	var zeroValue LcpResponse[[]SteamGame]
	games, err := fetchCache[[]SteamGame]("steam")
	if err != nil {
		return zeroValue, err
	}
	return games, nil
}
