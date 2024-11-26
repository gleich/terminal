package lcp

type AppleMusicCacheData struct {
	RecentlyPlayed []AppleMusicSong `json:"recently_played"`
}

type AppleMusicSong struct {
	Track            string   `json:"track"`
	Artist           string   `json:"artist"`
	Album            string   `json:"album"`
	Genres           []string `json:"genres"`
	ReleaseDate      string   `json:"release_date"`
	DurationInMillis int      `json:"duration_in_millis"`
	AlbumArtURL      string   `json:"album_art_url"`
	URL              string   `json:"url"`
	ID               string   `json:"id"`
}

func FetchAppleMusicCache() (LcpResponse[AppleMusicCacheData], error) {
	var zeroValue LcpResponse[AppleMusicCacheData]
	data, err := fetchCache[AppleMusicCacheData]("applemusic")
	if err != nil {
		return zeroValue, nil
	}
	return data, nil
}
