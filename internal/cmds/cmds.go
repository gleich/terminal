package cmds

import (
	"io"
	"strings"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/lcp/pkg/lcp"
	"go.mattglei.ch/terminal/internal/output"
	"go.mattglei.ch/timber"
	"golang.org/x/term"
)

func Terminal(s ssh.Session, styles output.Styles, client *lcp.Client) {
	prefix := styles.Green.Render("λ ")
	terminal := term.NewTerminal(s, prefix)

	for {
		cmd, err := terminal.ReadLine()
		if err == io.EOF {
			output.Line(s)
			return
		}
		if err != nil {
			timber.Error(err, "failed to process command")
			output.Line(s, "processing command failed, closing connection")
			return
		}

		switch strings.ToLower(strings.Trim(cmd, " ")) {
		case "":
		case "exit":
			return
		case "help":
			output.Line(s, output.Help(styles))
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
			output.Linef(
				s,
				"\nInvalid command '%s'. Type 'help' to see available commands.\n\n",
				cmd,
			)
		}
	}
}
