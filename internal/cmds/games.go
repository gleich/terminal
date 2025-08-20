package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/lcp/pkg/lcp"
	"go.mattglei.ch/terminal/internal/output"
	"go.mattglei.ch/terminal/internal/util"
	"go.mattglei.ch/timber"
)

func games(s ssh.Session, styles output.Styles, client *lcp.Client) {
	var (
		headers = []string{"", "NAME", "ACHIEVEMENT PROGRESS", "TIME IN GAME", "STEAM LINK"}
		data    [][]string
	)
	games, err := lcp.FetchCache[[]lcp.SteamGame](client)
	if err != nil {
		msg := "failed to load steam games from lcp"
		output.Line(s, styles.Red.Render(msg))
		timber.Error(err, msg)
		return
	}

	for i, g := range games.Data {
		var achievementProgress string
		if g.AchievementProgress == nil {
			achievementProgress = "N/A"
		} else {
			achievementProgress = fmt.Sprintf("%.2f%%", *g.AchievementProgress)
		}
		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				g.Name,
				achievementProgress,
				util.RenderDuration(int(g.PlaytimeForever * 60)),
				g.URL,
			},
		)
	}

	output.Line(
		s,
		"\nTo relax I love to play games with some of my friends. Below are some of my most recent games from Steam:",
	)

	output.Line(s)
	table := output.Table(styles).Headers(headers...).Rows(data...).Render()
	output.Line(s, table)
	output.LiveFrom(s, styles, table, games.Updated)
	output.Line(s)
}
