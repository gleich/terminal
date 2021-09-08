package util

import (
	"fmt"
	"time"

	"github.com/gliderlabs/ssh"
)

// Output a typewriter animation on a string where each character slowly gets outputted
func TypewriterAnimation(s ssh.Session, str string) {
	for _, char := range str {
		fmt.Fprint(s, string(char))
		time.Sleep(time.Millisecond * 20)
	}
}
