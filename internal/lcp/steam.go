package lcp

import (
	"pkg.mattglei.ch/lcp-2/pkg/models"
)

func FetchGames() (response[[]models.SteamGame], error) {
	var zeroValue response[[]models.SteamGame]
	games, err := fetchCache[[]models.SteamGame]("steam")
	if err != nil {
		return zeroValue, err
	}
	return games, nil
}
