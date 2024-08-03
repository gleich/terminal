package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/terminal/internal/cmds"
	"github.com/gleich/terminal/internal/output"
	"github.com/joho/godotenv"
	"golang.org/x/term"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		lumber.Fatal(err, "loading .env failed")
	}

	go startHTTP()
	startSSH()
}

func startHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/gleich/terminal", http.StatusTemporaryRedirect)
	})

	lumber.Info("starting http server")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		lumber.Fatal(err, "starting http server failed")
	}
}

func startSSH() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		lumber.Fatal(err, "getting home directory failed")
	}

	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("0.0.0.0", "22")),
		wish.WithHostKeyPath(filepath.Join(homedir, ".ssh", "id_rsa")),
		wish.WithMiddleware(func(next ssh.Handler) ssh.Handler {
			return func(s ssh.Session) {
				out := output.OutputFromSession(s)
				colors := output.NewColors(out.ColorProfile())

				fmt.Fprintln(s)
				output.Typewriter(
					s,
					50*time.Millisecond,
					out.String("CONNECTION SUCCESSFULLY ESTABLISHED TO TERMINAL").
						Bold().
						Underline().
						String(),
				)
				output.Typewriter(
					s,
					30*time.Millisecond,
					out.String("\nWelcome to my personal terminal! Enter `help` to available commands.\n").
						Foreground(colors.Green).
						String(),
				)

				prefix := out.String("λ ").Foreground(colors.Green)
				terminal := term.NewTerminal(s, prefix.String())
				consecutiveFails := 0
				for {
					cmd, err := terminal.ReadLine()
					if err == io.EOF {
						fmt.Fprintln(s)
						return
					}
					if err != nil {
						fmt.Println(err.Error())
						lumber.Error(err, "processing new command failed")
					}
					switch cmd {
					case "":
					case "exit":
						return
					case "help":
						cmds.Help(s)
					case "workouts":
						cmds.Workouts(s, out, colors)
					case "clear":
						out.ClearScreen()
					default:
						fmt.Fprintf(s, "\nInvalid command '%s'.\n\n", cmd)
						consecutiveFails++
						if consecutiveFails > 10 {
							return
						}
					}
				}
			}
		},
			logging.Middleware(), activeterm.Middleware()),
	)
	if err != nil {
		lumber.Fatal(err, "creating server failed")
	}

	lumber.Info("starting ssh server")
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		lumber.Fatal(err, "starting server failed")
	}
}
