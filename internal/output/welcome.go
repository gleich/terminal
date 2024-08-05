package output

import (
	"fmt"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/muesli/termenv"
)

func Welcome(s ssh.Session, out *termenv.Output, colors Colors) {
	fmt.Fprintln(s)
	Typewriter(
		s,
		50*time.Millisecond,
		out.String("CONNECTION SUCCESSFULLY ESTABLISHED TO TERMINAL").
			Bold().
			Underline().
			String(),
	)
	Typewriter(
		s,
		30*time.Millisecond,
		out.String("\nWelcome to Matt Gleich's terminal. Enter `help` to see available commands.\n").
			Foreground(colors.Green).
			String(),
	)
}
