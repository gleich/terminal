package cmds

import (
	"fmt"
	"io"

	"github.com/charmbracelet/ssh"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/terminal/internal/output"
	"golang.org/x/term"
)

func Terminal(s ssh.Session, colors output.Styles) {
	prefix := colors.Green.Render("λ ")
	terminal := term.NewTerminal(s, prefix)

	for {
		cmd, err := terminal.ReadLine()
		if err == io.EOF {
			fmt.Fprintln(s)
			return
		}
		if err != nil {
			lumber.Error(err, "failed to process command")
			fmt.Fprintln(s, "processing command failed, closing connection")
			return
		}

		switch cmd {
		case "":
		case "exit":
			return
		case "help":
			fmt.Fprintln(s, output.Help(colors))
		case "clear", "c":
			colors.Renderer.Output().ClearScreen()
		case "workouts":
			workouts(s, colors)
		default:
			fmt.Fprintf(s, "\nInvalid command '%s'.\n\n", cmd)
		}
	}
}
