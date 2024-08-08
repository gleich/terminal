package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"github.com/gleich/terminal/internal/lcp"
	"github.com/gleich/terminal/internal/output"
	"github.com/gleich/terminal/internal/util"
)

func workouts(s ssh.Session, styles output.Styles) {
	headers := []string{"", "NAME", "STARTED", "DURATION", "DISTANCE", "TYPE", "AVG. HEART RATE"}

	var data [][]string
	activities, err := lcp.FetchActivities()
	if err != nil {
		fmt.Fprintln(s, styles.Red.Render("failed to load strava activities from lcp"))
		return
	}
	for i, a := range activities.Data {
		switch a.SportType {
		case "GravelRide":
			a.SportType = "Gravel Ride"
		case "MountainBikeRide":
			a.SportType = "Mountain Bike Ride"
		case "":
			a.SportType = "Workout"
		}
		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				a.Name,
				util.RenderExactFromNow(a.StartDate),
				util.RenderDuration(int(a.MovingTime)),
				fmt.Sprintf("%.2f mi [%.2f km]", a.Distance*0.000621371, a.Distance*0.001),
				a.SportType,
				fmt.Sprintf("%.2f bpm", a.AverageHeartrate),
			},
		)
	}

	fmt.Fprintln(s, output.Table(styles).Headers(headers...).Rows(data...))
}
