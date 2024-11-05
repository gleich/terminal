package output

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Table(styles Styles) *table.Table {
	var (
		headerStyle = styles.Green.Bold(true).Align(lipgloss.Center).Padding(0, 2)
		baseStyle   = styles.Renderer.NewStyle().Padding(0, 1)
		indexStyle  = styles.Grey.Align(lipgloss.Center)
		nameStyle   = styles.Renderer.NewStyle().Bold(true).Padding(0, 1).Align(lipgloss.Center)
	)
	return table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(styles.Grey).
		BorderRow(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == -1 {
				return headerStyle
			}
			if col == 0 {
				return indexStyle
			}
			if row != -1 && col == 1 {
				return nameStyle
			}
			return baseStyle
		})
}
