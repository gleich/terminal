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
			"\nI'm a third-year Computer Science student at the Rochester Institute of Technology (RIT). This website serves as a portfolio showcasing some of my projects, work experience, and personal interests. I am committed to applying my technical skills to develop practical solutions and consistently strive to expand my expertise. If you'd like to discuss a project or explore potential opportunities, please contact me at mail@mattglei.ch. Additional details about my professional work are in my résumé [%s].\n",
			linkStyle.Render("https://mattglei.ch/resume.pdf"),
		)),
	)
}
