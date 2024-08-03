package cmds

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/ssh"
	"github.com/gleich/terminal/internal/lcp"
	"github.com/gleich/terminal/internal/output"
	"github.com/gleich/terminal/internal/util"
	"github.com/muesli/termenv"
)

func Workouts(s ssh.Session, out *termenv.Output, colors output.Colors) {
	response, err := lcp.FetchActivities()
	if err != nil {
		fmt.Fprintln(s, "Fetching workouts from lcp failed.")
	}

	fmt.Fprintln(s)
	fmt.Fprintf(
		s,
		"Here are my 3 most recent workouts from Strava. Last updated %s from https://mattglei.ch/lcp.\n",
		util.RenderExactFromNow(response.Updated),
	)
	fmt.Fprintln(s)

	for i, a := range response.Data {
		t := table.New()
		t.Border(lipgloss.NormalBorder())
		t.BorderRow(true)
		t.Row(
			out.String(" DISTANCE ").Foreground(colors.Green).Bold().String(),
			fmt.Sprintf(" %.2f mi [%.2f km]", a.Distance*0.000621371, a.Distance*0.001),
		)
		t.Row(
			out.String(" AVERAGE HR ").Foreground(colors.Green).Bold().String(),
			fmt.Sprintf(" %.2f bpm", a.AverageHeartrate),
		)
		t.Row(
			out.String(" STRAVA LINK ").Foreground(colors.Green).Bold().String(),
			" "+out.String(fmt.Sprintf("https://strava.com/activities/%d", a.ID)).
				Underline().
				String()+" ",
		)
		fmt.Fprintf(
			s,
			"#%d: %s %s\n",
			i+1,
			out.String(a.Name).Foreground(colors.Blue).Bold().Underline(),
			out.String("["+util.RenderExactFromNow(a.StartDate)+"]").Foreground(colors.Grey),
		)
		fmt.Fprintln(s, t.Render())
		fmt.Fprintln(s)
	}
}
