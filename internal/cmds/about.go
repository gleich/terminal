package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/terminal/internal/output"
)

func about(s ssh.Session, styles output.Styles) {
	linkStyle := styles.Blue.Underline(true)
	fmt.Fprintln(
		s,
		styles.Renderer.NewStyle().Width(output.MAX_WIDTH).Render(fmt.Sprintf(
			"\nHey! I'm Matt Gleich, a college student attending the Rochester Institute of Technology (RIT). I'm going into my 3rd year studying computer science there and in my free time I really enjoy cycling (gravel, road, & mountain bike) and photography. This summer I will be working at KCF Technologies as a DevOps intern. Want to get in touch? Feel free to shoot me an email over at email@mattglei.ch. More of my professional work is detailed in my résumé [%s].\n",
			linkStyle.Render("https://mattglei.ch/resume.pdf"),
		)),
	)
}
