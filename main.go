package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gleich/logoru"
	"github.com/gleich/ssh/pkg/colors"
	"github.com/gleich/ssh/pkg/commands"
	"github.com/gleich/ssh/pkg/messages"
	"github.com/gleich/ssh/pkg/web"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	logoru.Info("Started program")

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, web.HTML)
			logoru.Success("Handled http request")
		})

		err := http.ListenAndServe(os.Getenv("GLEICH_SSH_HTTP_PORT"), nil)
		if err != nil {
			logoru.Critical("Failed to start http server", err)
		}
		logoru.Info("Started http server")
	}()

	ssh.Handle(func(s ssh.Session) {
		logoru.Info("Handling session")
		messages.OutputWelcome(s)

		terminal := term.NewTerminal(s, colors.Green.Sprint("λ "))
		consecutive_fails := 0
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
			case "turtles":
				commands.RunTurtle(s)
			default:
				fmt.Fprintln(s, colors.Red.Sprint("Please enter a valid command"))
				consecutive_fails++
				if consecutive_fails > 10 {
					fmt.Fprintln(s, colors.Red.Sprint("Pipe broke to prevent attack"))
					break
				}
			}
		}
	})

	err := ssh.ListenAndServe(os.Getenv("GLEICH_SSH_PORT"), nil)
	if err != nil {
		logoru.Critical("Failed to start ssh server", err)
	}
	logoru.Info("Started ssh server")
}
