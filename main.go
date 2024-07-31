package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/ssh/internal/util"
	"github.com/gliderlabs/ssh"
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
		util.WriteToConnection(s, "Hello world!")
	})

	lumber.Info("starting ssh server")

	homedir, err := os.UserHomeDir()
	if err != nil {
		lumber.Fatal(err, "getting home directory failed")
	}

	err = ssh.ListenAndServe(":22", nil, ssh.HostKeyFile(filepath.Join(homedir, ".ssh", "id_rsa")))
	if err != nil {
		lumber.Fatal(err, "starting server failed")
	}
}
