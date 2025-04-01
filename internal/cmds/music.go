package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/lcp-2/pkg/lcp"
	"go.mattglei.ch/terminal/internal/output"
	"go.mattglei.ch/timber"
)

func music(s ssh.Session, styles output.Styles, client *lcp.Client) {
	cacheData, err := lcp.FetchCache[lcp.AppleMusicCache](client)
	if err != nil {
		msg := "failed to load data from apple music cache"
		fmt.Fprintln(s, styles.Red.Render(msg))
		timber.Error(err, msg)
		return
	}

	fmt.Fprintln(s)
	fmt.Fprintln(
		s,
		styles.Renderer.NewStyle().
			Width(output.MAX_WIDTH).
			Render("I love a lot of different types of music ranging from electronic to jazz. A few of my favorite artists are Daft Punk, The Smiths, Eagles, Mac DeMarco, Fleetwood Mac, Oasis, and Deftones."),
	)

	fmt.Fprintln(s, "\nHere are 5 of my most recently played songs from Apple Music:")

	for i, song := range cacheData.Data.RecentlyPlayed[5:] {
		fmt.Fprintf(
			s,
			"  %d. %s by %s\n",
			i+1,
			styles.Green.Bold(true).Render(song.Track),
			song.Artist,
		)
	}
	fmt.Fprintln(s)
}
