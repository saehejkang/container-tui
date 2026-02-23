package components

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderHelpOverlay(width, height int) string {
	content := lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Render("Keyboard Shortcuts"),
		"",
		"↑ / k       Move up",
		"↓ / j       Move down",
		"?           Toggle this help",
		"Esc         Close help",
		"q / Ctrl+C  Quit",
	)

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#0A84FF")).
		Background(lipgloss.Color("#1F1F1F")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(1, 3).
		Render(content)

	return box
}
