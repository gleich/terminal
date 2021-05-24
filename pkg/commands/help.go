package commands

import (
	"fmt"
	"strings"

	"github.com/gliderlabs/ssh"
)

// The help message listing all commands
const HelpMessage = `	help     Output this list of commands
	exit     Leave :(`

// Run the help command
func RunHelp(s ssh.Session) {
	fmt.Fprintln(s, "\n"+strings.ReplaceAll(HelpMessage, "	", "")+"\n")
}
