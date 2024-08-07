package cmds

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/ssh"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/terminal/internal/lcp"
	"github.com/gleich/terminal/internal/output"
	"github.com/gleich/terminal/internal/util"
)

func workouts(s ssh.Session, styles output.Styles) {
	headers := []string{"", "NAME", "START TIME", "DURATION", "DISTANCE", "TYPE", "AVG. HEART RATE"}

	var data [][]string
	activities, err := lcp.FetchActivities()
	if err != nil {
		fmt.Fprintln(s, styles.Red.Render("failed to load strava activities"))
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
		loc, err := time.LoadLocation(strings.Split(a.Timezone, " ")[1])
		if err != nil {
			lumber.Error(err, "loading timezone", a.Timezone, "failed")
		}

		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				a.Name,
				a.StartDate.In(loc).Format("01/02/2006 @ 03:04 PM MST"),
				util.RenderDuration(int(a.MovingTime)),
				fmt.Sprintf("%.2f mi [%.2f km]", a.Distance*0.000621371, a.Distance*0.001),
				a.SportType,
				fmt.Sprintf("%.2f bpm", a.AverageHeartrate),
			},
		)
	}

	var (
		headerStyle = styles.Green.Bold(true).Align(lipgloss.Center).Padding(0, 2)
		baseStyle   = styles.Renderer.NewStyle().Padding(0, 1)
		indexStyle  = styles.Grey.Align(lipgloss.Center)
		nameStyle   = styles.Renderer.NewStyle().Bold(true).Padding(0, 1).Align(lipgloss.Center)
	)
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(styles.Grey).
		BorderRow(true).
		Headers(headers...).
		Rows(data...).StyleFunc(func(row, col int) lipgloss.Style {
		if row == 0 {
			return headerStyle
		}
		if col == 0 {
			return indexStyle
		}
		if row != 0 && col == 1 {
			return nameStyle
		}
		return baseStyle
	})

	fmt.Fprintln(s, t)
}
