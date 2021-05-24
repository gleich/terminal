package main

import (
	"fmt"

	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/ssh_me/pkg/commands"
	"github.com/Matt-Gleich/ssh_me/pkg/messages"
	"github.com/fatih/color"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	logoru.Info("Started program")

	ssh.Handle(func(s ssh.Session) {
		logoru.Info("Handling session")
		messages.OutputWelcome(s)

		terminal := term.NewTerminal(s, "λ ")
		for {
			cmd, err := terminal.ReadLine()
			if err != nil {
				logoru.Error("Failed to process new command", err)
			}

			switch cmd {
			case "help":
				commands.RunHelp(s)
			case "exit":
				return
			default:
				fmt.Fprintln(s, color.RedString("Please enter a valid command"))
			}
		}
	})

	err := ssh.ListenAndServe(":2222", nil)
	if err != nil {
		logoru.Critical("Failed to start ssh server", err)
	}
}
