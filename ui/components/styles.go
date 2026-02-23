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

// Status colors for container rows.
var StatusRunningColor = lipgloss.Color("#32D74B") // green
var StatusStoppedColor = lipgloss.Color("#888888") // gray
var StatusErrorColor = lipgloss.Color("#FF3B30")   // red

// StatusColor maps a normalized status string to the correct Lip Gloss color.
func StatusColor(status string) lipgloss.Color {
	switch status {
	case "running":
		return StatusRunningColor
	case "error":
		return StatusErrorColor
	default:
		return StatusStoppedColor
	}
}

// Resource bar styles — filled and empty portions of CPU/memory bars.
var ResourceBarFilledStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#0A84FF")) // blue, matches CursorStyle

var ResourceBarEmptyStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#333333"))

// Container row styles for the left panel.
var ContainerRowStyle = lipgloss.NewStyle().
	Padding(0, 1)

var ContainerRowSelectedStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#0A84FF")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Padding(0, 1)
