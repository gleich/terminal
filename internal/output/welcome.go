package output

import (
	"fmt"
	"time"

	"github.com/charmbracelet/ssh"
)

func Welcome(s ssh.Session, colors Styles) {
	TypewriterAnimation(s, 60*time.Millisecond, "\nESTABLISHING CONNECTION")
	TypewriterAnimation(s, 500*time.Millisecond, " ...")
	TypewriterAnimation(
		s,
		50*time.Millisecond,
		" CONNECTION SUCCESSFULLY ESTABLISHED",
	)
	fmt.Fprintln(s)
	fmt.Fprintln(s)
	TypewriterAnimation(
		s,
		50*time.Millisecond,
		colors.Green.Render("Welcome to Matt Gleich's personal terminal."),
	)
	fmt.Fprintln(s)
	fmt.Fprintln(s, Help(colors))
}
