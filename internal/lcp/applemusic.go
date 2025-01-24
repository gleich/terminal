package lcp

import "pkg.mattglei.ch/lcp-2/pkg/models"

func FetchAppleMusicCache() (response[models.AppleMusicCache], error) {
	var zeroValue response[models.AppleMusicCache]
	data, err := fetchCache[models.AppleMusicCache]("applemusic")
	if err != nil {
		return zeroValue, nil
	}
	return data, nil
}
