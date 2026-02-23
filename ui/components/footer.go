package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderFooter(text string, width int) string {
	line := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render(strings.Repeat("─", width))

	return FooterStyle.Render(text) + "\n" + line
}
