package commands

import (
	"fmt"
	"strings"

	"github.com/gleich/ssh/pkg/util"
	"github.com/gliderlabs/ssh"
)

// Run the turtle command
func RunTurtle(s ssh.Session) {
	turtles := []string{}
	for i := 0; i < 100; i++ {
		turtles = append(turtles, "🐢")
	}
	util.TypewriterAnimation(s, strings.Join(turtles, " "))
	fmt.Fprintln(s)
}
