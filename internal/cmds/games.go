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

	for i, g := range games.Data[:5] {
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

	fmt.Fprintln(s, output.Table(styles).Headers(headers...).Rows(data...))
}
