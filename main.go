package main

import (
	"fmt"
	"io"

	"github.com/gleich/lumber/v2"
	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		lumber.Info("handling connection from", s.RemoteAddr().String())
		_, err := io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		if err != nil {
			lumber.Error(err, "failed to write string to user")
		}
	})

	lumber.Info("starting server")
	lumber.Fatal(ssh.ListenAndServe(":22", nil))
}
