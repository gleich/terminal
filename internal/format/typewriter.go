package format

import (
	"fmt"
	"time"

	"github.com/gliderlabs/ssh"
)

func OutputTypewriter(s ssh.Session, speed time.Duration, msg string) {
	for _, c := range msg {
		fmt.Fprint(s, string(c))
		time.Sleep(speed)
	}
	fmt.Fprintln(s)
}
