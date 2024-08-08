package cmds

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/gleich/terminal/internal/lcp"
	"github.com/gleich/terminal/internal/output"
	"github.com/gleich/terminal/internal/util"
)

func projects(s ssh.Session, styles output.Styles) {
	headers := []string{"", "NAME", "DESCRIPTION", "UPDATED", "LANGUAGE", "GITHUB LINK"}
	var data [][]string
	repositories, err := lcp.FetchRepositories()
	if err != nil {
		fmt.Fprintln(s, styles.Red.Render("failed to load github repositories from lcp"))
		return
	}
	for i, p := range repositories.Data {
		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				p.Name,
				p.Description,
				util.RenderExactFromNow(p.UpdatedAt),
				fmt.Sprintf(
					"%s %s",
					styles.Renderer.NewStyle().
						Foreground(lipgloss.Color(p.LanguageColor)).
						Render("●"),
					p.Language,
				),
				p.URL,
			},
		)
	}

	fmt.Fprintln(s, output.Table(styles).Headers(headers...).Rows(data...))
}
