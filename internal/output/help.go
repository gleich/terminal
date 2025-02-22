package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Help(styles Styles) string {
	box := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(styles.Grey.GetForeground()).
		Padding(1, 2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Margin(1).
		MarginRight(0)
	cmdStyle := styles.Green.Bold(true)
	return box.Render(fmt.Sprintf(`%s     a little about myself
%s  my recent workouts from Strava
%s  recent projects I've worked on from GitHub
%s     games I've recently played on Steam
%s     music I've listened to recently

%s   displays this help table
%s   exit out of terminal
%s  clear the terminal`,
		cmdStyle.Render("about"),
		cmdStyle.Render("workouts"),
		cmdStyle.Render("projects"),
		cmdStyle.Render("games"),
		cmdStyle.Render("music"),
		cmdStyle.Render("help"),
		cmdStyle.Render("exit"),
		cmdStyle.Render("clear"),
	))
}
