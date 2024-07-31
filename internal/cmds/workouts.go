package cmds

import (
	"fmt"

	"github.com/gleich/ssh/internal/format"
	"github.com/gleich/ssh/internal/lcp"
	"github.com/gleich/ssh/internal/util"
	"github.com/gliderlabs/ssh"
)

func Workouts(s ssh.Session) {
	response, err := lcp.FetchActivities()
	if err != nil {
		fmt.Fprintln(s, "Fetching workouts from lcp failed.")
	}

	fmt.Fprintln(s)
	fmt.Fprintln(s, "Here are my 3 most recent workouts from Strava. Last updated "+util.RenderExactFromNow(response.Updated))
	fmt.Fprintln(s)
	for i, a := range response.Data[:3] {
		fmt.Fprintf(s, "#%d: %s %s\n", i+1, format.UnderlinedBold(a.Name), format.Grey.Sprint("["+util.RenderExactFromNow(a.StartDate)+"]"))
		fmt.Fprintln(s)
		fmt.Fprintf(s, "\tDistance: %.2f mi [%.2f km]\n", a.Distance*0.000621371, a.Distance*0.001)
		fmt.Fprintf(s, "\tAverage heart rate: %.2f bpm\n", a.AverageHeartrate)
		fmt.Fprintf(s, "\tStrava activity: https://strava.com/activities/%d\n", a.ID)
		fmt.Fprintln(s)
	}
}
