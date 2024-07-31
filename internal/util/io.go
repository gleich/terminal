package util

import (
	"io"

	"github.com/gleich/lumber/v2"
	"github.com/gliderlabs/ssh"
)

func WriteToConnection(connection ssh.Session, msg string) {
	_, err := io.WriteString(connection, msg)
	if err != nil {
		lumber.Error(err, "failed to write", msg, "to connection")
	}
}
