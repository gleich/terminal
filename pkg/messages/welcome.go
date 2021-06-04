package messages

import (
	"fmt"
	"time"

	"github.com/gleich/ssh_me/pkg/colors"
	"github.com/gleich/ssh_me/pkg/commands"
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
	for _, char := range message {
		fmt.Fprint(s, string(char))
		time.Sleep(time.Millisecond * 20)
	}
}
