package main

import (
	"io"
	"log"

	"github.com/Matt-Gleich/logoru"
	"github.com/gliderlabs/ssh"
)

func main() {
	logoru.Info("Started program")

	ssh.Handle(func(s ssh.Session) {
		logoru.Info("Handling session")
		io.WriteString(s, "Hello World!\n")
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
