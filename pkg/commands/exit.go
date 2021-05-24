package commands

import (
	"fmt"

	"github.com/gliderlabs/ssh"
)

// Run the exit command
func RunExit(s ssh.Session) {
	fmt.Fprintln(s, "\nSorry to see you go! Have a good day :)")
}
