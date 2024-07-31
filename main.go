package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/ssh/internal/cmds"
	"github.com/gleich/ssh/internal/colors"
	"github.com/gleich/ssh/internal/util"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	go startHTTP()
	startSSH()
}

func startHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/gleich/ssh", http.StatusTemporaryRedirect)
		lumber.Success("handled http redirect to github repo")
	})

	lumber.Info("starting http server")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		lumber.Fatal(err, "starting http server failed")
	}
}

func startSSH() {
	ssh.Handle(func(s ssh.Session) {
		lumber.Info("handling connection from", s.RemoteAddr().String())

		util.WriteToConnection(s, colors.Green.Sprint("Welcome to Matt Gleich's personal terminal! Enter `help` to available commands.\n"))

		terminal := term.NewTerminal(s, colors.Green.Sprint("λ "))
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
			default:
				fmt.Fprintf(s, "\nInvalid command '%s'.\n\n", cmd)
				consecutiveFails++
				if consecutiveFails > 10 {
					break
				}
			}
		}
	})

	homedir, err := os.UserHomeDir()
	if err != nil {
		lumber.Fatal(err, "getting home directory failed")
	}

	lumber.Info("starting ssh server")
	err = ssh.ListenAndServe(":22", nil, ssh.HostKeyFile(filepath.Join(homedir, ".ssh", "id_rsa")))
	if err != nil {
		lumber.Fatal(err, "starting server failed")
	}
}
