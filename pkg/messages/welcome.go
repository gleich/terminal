package messages

import (
	"fmt"

	"github.com/gleich/ssh/pkg/colors"
	"github.com/gleich/ssh/pkg/commands"
	"github.com/gleich/ssh/pkg/util"
	"github.com/gliderlabs/ssh"
)

// Output a welcome message to the user
func OutputWelcome(s ssh.Session) {
	fmt.Fprintln(s, colors.Red.Sprint(`
┌┬┐┌─┐┌┬┐┌┬┐  ┌─┐┬  ┌─┐┬┌─┐┬ ┬
│││├─┤ │  │   │ ┬│  ├┤ ││  ├─┤
┴ ┴┴ ┴ ┴  ┴   └─┘┴─┘└─┘┴└─┘┴ ┴
`))

	message := "👋 Hello and welcome to my ssh server!\nThis server acts like a terminal but with the following commands:\n\n" + commands.HelpMessage + "\n\n"
	util.TypewriterAnimation(s, message)
}
