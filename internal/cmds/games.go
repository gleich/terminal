package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"github.com/gleich/terminal/internal/lcp"
	"github.com/gleich/terminal/internal/output"
	"github.com/gleich/terminal/internal/util"
)

func games(s ssh.Session, styles output.Styles) {
	var (
		headers = []string{"", "NAME", "ACHIEVEMENT PROGRESS", "TIME IN GAME", "STEAM LINK"}
		data    [][]string
	)
	games, err := lcp.FetchGames()
	if err != nil {
		fmt.Fprintln(s, styles.Red.Render("failed to load steam games from lcp"))
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

	fmt.Fprintln(
		s,
		"\nTo relax I love to play games with some of my friends. Below are some of my most recent games from Steam:",
	)
	fmt.Fprintln(s)
	table := output.Table(styles).Headers(headers...).Rows(data...).Render()
	fmt.Fprintln(s, table)
	output.LiveFrom(s, styles, table, games.Updated)
	fmt.Fprintln(s)
}
