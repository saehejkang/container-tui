package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func StatusBox(isRunning bool) string {
	color := "#FF3B30" // red
	text := "STOPPED"
	if isRunning {
		color = "#32D74B" // green
		text = "RUNNING"
	}

	box := lipgloss.NewStyle().
		Background(lipgloss.Color(color)).
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 2).
		MarginLeft(1).
		Bold(true).
		Render(text)

	return box
}

func RenderHeaderWithStatus(title string, isRunning bool) string {
	line := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render(strings.Repeat("─", lipgloss.Width(title)+8)) // extra width

	header := HeaderStyle.Copy().
		Background(lipgloss.Color("#1F1F1F")).
		Padding(2, 4).
		Render(title)

	status := StatusBox(isRunning)

	return lipgloss.JoinHorizontal(lipgloss.Center, header, status) + "\n" + line
}
