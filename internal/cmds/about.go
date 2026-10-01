package cmds

import (
	"fmt"

	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/terminal/internal/output"
)

func about(s ssh.Session, styles output.Styles) {
	linkStyle := styles.Blue.Underline(true)

	output.Line(s, "\nHey, I'm Matt Gleich")
	output.Line(
		s,
		styles.Renderer.NewStyle().Width(output.MAX_WIDTH).Render(fmt.Sprintf(
			"I'm a fifth-year computer science student (5 year program) at the Rochester Institute of Technology [%s] and currently a Software Engineer at Synthient [%s]. This website serves as a portfolio showcasing some of my projects, work experience, and personal interests. I am committed to applying my technical skills to develop practical solutions and consistently strive to expand my expertise. If you'd like to discuss a project or explore potential opportunities, please contact me at %s. Additional details about my professional work are in my résumé [%s].\n",
			linkStyle.Render("https://rit.edu"),
			linkStyle.Render("https://synthient.com"),
			linkStyle.Render("email@mattglei.ch"),
			linkStyle.Render("https://mattglei.ch/resume"),
		)),
	)
}
