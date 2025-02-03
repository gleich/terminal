package cmds

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/ssh"
	"golang.org/x/term"
	"pkg.mattglei.ch/lcp-2/pkg/lcp"
	"pkg.mattglei.ch/terminal/internal/output"
	"pkg.mattglei.ch/timber"
)

func Terminal(s ssh.Session, styles output.Styles, client *lcp.Client) {
	prefix := styles.Green.Render("λ ")
	terminal := term.NewTerminal(s, prefix)

	for {
		cmd, err := terminal.ReadLine()
		if err == io.EOF {
			fmt.Fprintln(s)
			return
		}
		if err != nil {
			timber.Error(err, "failed to process command")
			fmt.Fprintln(s, "processing command failed, closing connection")
			return
		}

		switch strings.ToLower(strings.Trim(cmd, " ")) {
		case "":
		case "exit":
			return
		case "help":
			fmt.Fprintln(s, output.Help(styles))
		case "clear", "c":
			styles.Renderer.Output().ClearScreen()
		case "about":
			about(s, styles)
		case "workouts":
			workouts(s, styles, client)
		case "projects":
			projects(s, styles, client)
		case "games":
			games(s, styles, client)
		case "music":
			music(s, styles, client)
		default:
			fmt.Fprintf(
				s,
				"\nInvalid command '%s'. Type `help` to see available commands.\n\n",
				cmd,
			)
		}
	}
}
