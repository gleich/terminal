package main

import (
	"net/http"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/ssh/internal/util"
	"github.com/gliderlabs/ssh"
)

func main() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://github.com/gleich/ssh", http.StatusTemporaryRedirect)
			lumber.Success("handled http redirect to github repo")
		})

		lumber.Info("starting http server")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			lumber.Fatal(err, "Failed to start http server")
		}
	}()

	ssh.Handle(func(s ssh.Session) {
		lumber.Info("handling connection from", s.RemoteAddr().String())
		util.WriteToConnection(s, "Hello world!")
	})

	lumber.Info("starting ssh server")
	err := ssh.ListenAndServe(":22", nil)
	if err != nil {
		lumber.Fatal(err, "starting server failed")
	}
}
