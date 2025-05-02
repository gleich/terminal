package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/lcp/pkg/lcp"
	"go.mattglei.ch/terminal/internal/output"
	"go.mattglei.ch/terminal/internal/util"
	"go.mattglei.ch/timber"
)

func workouts(s ssh.Session, styles output.Styles, client *lcp.Client) {
	activities, err := lcp.FetchCache[[]lcp.Workout](client)
	if err != nil {
		msg := "failed to load workouts from lcp"
		fmt.Fprintln(s, styles.Red.Render(msg))
		timber.Error(err, msg)
		return
	}

	fmt.Fprintln(
		s,
		"\nOne of my favorite things in the world is staying active and enjoying the outdoors. I grew up in New Hampshire hiking, biking, snowshoeing, and traveling with my family. Out of all of those things I especially love cycling mainly through gravel cycling, road cycling, and mountain biking. Recently I've been getting into lifting which has been a ton of fun. Below are 5 of my most recent Strava and Hevy workouts:",
	)
	for i, a := range activities.Data[:3] {
		switch a.SportType {
		case "GravelRide":
			a.SportType = "Gravel Ride"
		case "MountainBikeRide":
			a.SportType = "Mountain Bike Ride"
		case "WeightTraining":
			a.SportType = "Weight Training"
		case "":
			a.SportType = "Workout"
		}

		fmt.Fprintln(s)
		fmt.Fprintf(
			s,
			"%d. %s %s\n",
			i+1,
			styles.Green.Bold(true).Render(a.Name),
			styles.Grey.Render(fmt.Sprintf("[%s]", a.Platform)),
		)
		fmt.Fprintf(s, "    Started: %s\n", util.RenderExactFromNow(a.StartDate))
		fmt.Fprintf(s, "    Duration: %s\n", util.RenderDuration(int(a.MovingTime)))
		if a.Platform == "strava" {
			fmt.Fprintf(s, "    Type: %s\n", a.SportType)
			if a.Distance != 0 {
				fmt.Fprintf(
					s,
					"    Distance: %s\n",
					fmt.Sprintf("%.2f mi [%.2f km]", a.Distance*0.000621371, a.Distance*0.001),
				)
			}
			if a.AverageHeartrate != 0 {
				fmt.Fprintf(
					s,
					"    Avg. Heartrate: %s\n",
					fmt.Sprintf("%.2f bpm", a.AverageHeartrate),
				)
			}
		} else {
			fmt.Fprintln(s, "    Exercises:")
			for i, exercise := range a.HevyExercises {
				fmt.Fprintf(s, "        %s (%d/%d)\n", styles.Blue.Render(exercise.Title), i+1, len(a.HevyExercises))
				for j, set := range exercise.Sets {
					fmt.Fprintf(s, "            [set %d] %.1f lbs × %d reps\n", j+1, set.WeightKg*2.2046226218, set.Reps)
				}
			}
		}
	}
}
