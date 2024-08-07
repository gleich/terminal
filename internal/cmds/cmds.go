package cmds

import (
	"fmt"
	"io"

	"github.com/charmbracelet/ssh"
	"github.com/gleich/lumber/v2"
	"golang.org/x/term"
)

func Terminal(s ssh.Session) {
	prefix := "λ "
	terminal := term.NewTerminal(s, prefix)

	for {
		cmd, err := terminal.ReadLine()
		if err == io.EOF {
			fmt.Fprintln(s)
			return
		}
		if err != nil {
			lumber.Error(err, "failed to process command")
			fmt.Fprintln(s, "processing command failed, closing connection")
			return
		}

		switch cmd {
		case "":
		case "exit":
			return
		case "help":
			fmt.Fprintln(s, "hello world!")
		default:
			fmt.Fprintf(s, "\nInvalid command '%s'.\n\n", cmd)
		}
	}
}
