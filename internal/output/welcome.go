package output

import (
	"fmt"
	"time"

	"github.com/charmbracelet/ssh"
)

func Welcome(s ssh.Session, colors Colors) {
	TypewriterAnimation(s, 60*time.Millisecond, "\nESTABLISHING CONNECTION")
	TypewriterAnimation(s, 500*time.Millisecond, " ...\n")
	TypewriterAnimation(
		s,
		50*time.Millisecond,
		"CONNECTION SUCCESSFULLY ESTABLISHED TO TERMINAL",
	)
	fmt.Fprintln(s)
	fmt.Fprintln(s)
	TypewriterAnimation(
		s,
		50*time.Millisecond,
		colors.Green.Render(
			"Welcome to Matt Gleich's personal terminal. Here are the available commands:",
		),
	)
	fmt.Fprintln(s)
	fmt.Fprintf(s, Help(colors))
	fmt.Fprintln(s)
	fmt.Fprintln(s)
}
