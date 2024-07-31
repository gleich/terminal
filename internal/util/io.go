package util

import (
	"io"
	"time"

	"github.com/gleich/lumber/v2"
	"github.com/gliderlabs/ssh"
)

func WriteToConnection(connection ssh.Session, msg string) {
	msg = msg + "\n"
	for _, c := range msg {
		_, err := io.WriteString(connection, string(c))
		if err != nil {
			lumber.Error(err, "failed to write", msg, "to connection")
		}
		time.Sleep(30 * time.Millisecond)
	}
}
