package lcp

type AppleMusicCacheData struct {
	RecentlyPlayed []AppleMusicSong `json:"recently_played"`
}

type AppleMusicSong struct {
	Track            string `json:"track"`
	Artist           string `json:"artist"`
	DurationInMillis int    `json:"duration_in_millis"`
	AlbumArtURL      string `json:"album_art_url"`
	URL              string `json:"url"`
	ID               string `json:"id"`
}

func FetchAppleMusicCache() (response[AppleMusicCacheData], error) {
	var zeroValue response[AppleMusicCacheData]
	data, err := fetchCache[AppleMusicCacheData]("applemusic")
	if err != nil {
		return zeroValue, nil
	}
	return data, nil
}
