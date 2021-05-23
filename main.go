package main

import (
	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/ssh_me/pkg/messages"
	"github.com/gliderlabs/ssh"
)

func main() {
	logoru.Info("Started program")

	ssh.Handle(func(s ssh.Session) {
		logoru.Info("Handling session")

		messages.FakeStep(s, "Hello world!")
	})

	err := ssh.ListenAndServe(":2222", nil)
	if err != nil {
		logoru.Critical("Failed to start ssh server", err)
	}
}
