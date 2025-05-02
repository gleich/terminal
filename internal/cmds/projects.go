package cmds

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"go.mattglei.ch/lcp/pkg/lcp"
	"go.mattglei.ch/terminal/internal/output"
	"go.mattglei.ch/terminal/internal/util"
	"go.mattglei.ch/timber"
)

func projects(s ssh.Session, styles output.Styles, client *lcp.Client) {
	var (
		headers = []string{"", "NAME", "DESCRIPTION", "UPDATED", "LANGUAGE", "GITHUB LINK"}
		data    [][]string
	)
	repositories, err := lcp.FetchCache[[]lcp.GitHubRepository](client)
	if err != nil {
		msg := "failed to load github repositories from lcp"
		fmt.Fprintln(s, styles.Red.Render(msg))
		timber.Error(err, msg)
		return
	}
	for i, p := range repositories.Data {
		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				fmt.Sprintf("%s/%s", p.Owner, p.Name),
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

	linkStyle := styles.Blue.Underline(true)
	table := output.Table(styles).Headers(headers...).Rows(data...).Render()
	fmt.Fprintln(
		s,
		styles.Renderer.NewStyle().Width(lipgloss.Width(table)+10).Render(fmt.Sprintf(
			"\nFor the past five years, I have been passionately pursuing programming. From developing PCBs with custom integrated circuit drivers written in Rust to creating CLIs and websites, I have explored various facets of the programming world. My journey includes cloud automation work at Bottomline Technologies [%s], where I utilized Python, Puppet, Docker, and Grafana. At Rootly [%s], I developed their official CLI in Golang. More recently, I contributed to Stainless API [%s] as an engineering developer, automating customer deployments and product testing.",
			linkStyle.Render("https://bottomline.com"),
			linkStyle.Render("https://rootly.com"),
			linkStyle.Render("https://stainlessapi.com"),
		)),
	)
	fmt.Fprintln(s, "\nHere are my pinned repositories from GitHub:")
	fmt.Fprintln(s)
	fmt.Fprintln(s, table)
	output.LiveFrom(s, styles, table, repositories.Updated)
	fmt.Fprintln(s)
}
