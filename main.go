package main

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/ssh_me/pkg/commands"
	"github.com/Matt-Gleich/ssh_me/pkg/messages"
	"github.com/gliderlabs/ssh"
	"github.com/wayneashleyberry/truecolor/pkg/color"
	"golang.org/x/term"
)

func main() {
	logoru.Info("Started program")

	ssh.Handle(func(s ssh.Session) {
		logoru.Info("Handling session")
		messages.OutputWelcome(s)

		terminal := term.NewTerminal(s, color.Color(0, 255, 0).Sprint("λ "))
		for {
			cmd, err := terminal.ReadLine()
			if err != nil {
				logoru.Error("Failed to process new command", err)
			}

			switch cmd {
			case "help":
				commands.RunHelp(s)
			case "exit":
				commands.RunExit(s)
				return
			default:
				fmt.Fprintln(s, color.Color(255, 0, 0).Sprint("Please enter a valid command"))
			}
		}
	})

	err := ssh.ListenAndServe(os.Getenv("SSH_ME_PORT"), nil)
	if err != nil {
		logoru.Critical("Failed to start ssh server", err)
	}
}
