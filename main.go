package main

import (
	"fmt"
	"os"

	"github.com/gleich/logoru"
	"github.com/gleich/ssh_me/pkg/colors"
	"github.com/gleich/ssh_me/pkg/commands"
	"github.com/gleich/ssh_me/pkg/messages"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	logoru.Info("Started program")

	ssh.Handle(func(s ssh.Session) {
		logoru.Info("Handling session")
		messages.OutputWelcome(s)

		terminal := term.NewTerminal(s, colors.Green.Sprint("λ "))
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
				fmt.Fprintln(s, colors.Red.Sprint("Please enter a valid command"))
			}
		}
	})

	err := ssh.ListenAndServe(os.Getenv("SSH_ME_PORT"), nil)
	if err != nil {
		logoru.Critical("Failed to start ssh server", err)
	}
}
