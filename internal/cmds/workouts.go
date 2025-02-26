package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/lcp-2/pkg/lcp"
	"go.mattglei.ch/terminal/internal/output"
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
		"\nOne of my favorite things in the world is staying active and enjoying the outdoors. I grew up in New Hampshire hiking, biking, snowshoeing, and traveling with my family. Out of all of those things I especially love cycling mainly through gravel cycling, road cycling, and mountain biking. Below are some of my most recent Strava activities:",
	)
	for i, a := range activities.Data {
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

		fmt.Fprintf(
			s,
			"%d. %s %s\n",
			i+1,
			styles.Green.Bold(true).Render(a.Name),
			styles.Grey.Render(fmt.Sprintf("[%s]", a.Platform)),
		)
	}
}
