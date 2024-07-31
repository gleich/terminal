package main

import (
	"github.com/gleich/lumber/v2"
	"github.com/gleich/ssh/internal/util"
	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		lumber.Info("handling connection from", s.RemoteAddr().String())
		util.WriteToConnection(s, "Hello world!")
	})

	lumber.Info("starting server")
	err := ssh.ListenAndServe(":22", nil)
	if err != nil {
		lumber.Fatal(err, "starting server failed")
	}
}
