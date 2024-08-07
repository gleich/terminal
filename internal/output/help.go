package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Help(colors Colors) string {
	box := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(colors.Grey.GetForeground()).
		Padding(0, 1).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
	cmdStyle := colors.Green.Underline(true)
	return box.Render(fmt.Sprintf(`%s  my recent workouts from Strava
%s  recent projects I've worked on from GitHub
%s     games I've recently played on Steam

%s   displays this help table
%s   exit out of terminal
%s  clear the terminal`,
		cmdStyle.Render("workouts"),
		cmdStyle.Render("projects"),
		cmdStyle.Render("games"),
		cmdStyle.Render("help"),
		cmdStyle.Render("exit"),
		cmdStyle.Render("clear"),
	))
}
