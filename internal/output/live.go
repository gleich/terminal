package output

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/gleich/terminal/internal/util"
)

func LiveFrom(s ssh.Session, styles Styles, table string, updated time.Time) {
	liveStyle := lipgloss.NewStyle().Width(lipgloss.Width(table)).Align(lipgloss.Center)
	fmt.Fprintln(
		s,
		styles.Red.Bold(true).Inherit(liveStyle).Render(
			"LIVE DATA FROM LCP (https://mattglei.ch/lcp)",
		),
	)
	fmt.Fprintln(
		s,
		liveStyle.Render(
			fmt.Sprintf("[Updated %s]", util.RenderExactFromNow(updated)),
		),
	)
}
