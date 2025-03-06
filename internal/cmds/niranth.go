package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/terminal/internal/output"
)

func niranth(s ssh.Session, styles output.Styles) {
	linkStyle := styles.Blue.Underline(true)
	fmt.Fprintln(
		s,
		styles.Renderer.NewStyle().Width(output.MAX_WIDTH).Render(fmt.Sprintf(
			"hey im niranth! okay bye[%s]",
			linkStyle.Render("https://instagram.com/niranthc"),
		)),
	)
}
