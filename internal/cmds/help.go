package cmds

import (
	"fmt"

	"github.com/gliderlabs/ssh"
)

func Help(s ssh.Session) {
	fmt.Fprintln(s, `
help      displays this help table
exit      exit out of terminal
workouts  my recent workouts of Strava`)
	fmt.Fprintln(s)
}
