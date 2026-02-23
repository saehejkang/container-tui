package components

import "github.com/charmbracelet/lipgloss"

var MenuStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#1F1F1F")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(1, 2)

var CursorStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#0A84FF")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1)

var OutputBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	MarginLeft(1)

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF5F5F")).
	Align(lipgloss.Center).
	Padding(1, 0)

var FooterStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#888888")).
	Padding(1, 2)
