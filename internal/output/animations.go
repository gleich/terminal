package output

import (
	"fmt"
	"time"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/timber"
)

func TypewriterAnimation(s ssh.Session, speed time.Duration, msg string) {
	for _, c := range msg {
		_, err := fmt.Fprint(s, string(c))
		if err != nil {
			timber.Error(err, "failed to output", c)
		}
		time.Sleep(speed)
	}
}
