package cmds

import (
	"fmt"

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
	fmt.Fprintf(s, "Here are my 3 most recent workouts from Strava. Last updated %s from https://mattglei.ch/lcp.\n", util.RenderExactFromNow(response.Updated))
	fmt.Fprintln(s)
	for i, a := range response.Data[:3] {
		fmt.Fprintf(s, "#%d: %s %s\n", i+1, out.String(a.Name).Foreground(colors.Blue).Bold().Underline(), out.String("["+util.RenderExactFromNow(a.StartDate)+"]").Foreground(colors.Grey))
		fmt.Fprintln(s)
		fmt.Fprintf(s, "\tDistance: %.2f mi [%.2f km]\n", a.Distance*0.000621371, a.Distance*0.001)
		fmt.Fprintf(s, "\tAverage heart rate: %.2f bpm\n", a.AverageHeartrate)
		fmt.Fprintf(s, "\tStrava activity: https://strava.com/activities/%d\n", a.ID)
		fmt.Fprintln(s)
	}
}
