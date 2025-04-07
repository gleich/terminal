package output

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/terminal/internal/util"
)

func LiveFrom(s ssh.Session, styles Styles, table string, updated time.Time) {
	liveStyle := lipgloss.NewStyle().Width(lipgloss.Width(table)).Align(lipgloss.Center)
	fmt.Fprintln(
		s,
		styles.Red.Bold(true).Inherit(liveStyle).Render(
			"LIVE DATA FROM LCP (https://mattglei.ch/writing/lcp)",
		),
	)
	fmt.Fprintln(
		s,
		liveStyle.Render(
			fmt.Sprintf("[Updated %s]", util.RenderExactFromNow(updated)),
		),
	)
}
