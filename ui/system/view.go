package system

import (
	"container-tui/ui/components"
	"container-tui/ui/system/subcommands"

	"github.com/charmbracelet/lipgloss"
)

func RenderSystem(m *Model) string {
	isRunning := false
	if statusModel, ok := m.ActiveView.(*subcommands.StatusModel); ok {
		status := statusModel.Fields["status"]
		if status == "running" || status == "started" {
			isRunning = true
		}
	}

	header := components.RenderHeaderWithStatus("Container TUI", isRunning)

	menuLines := make([]string, len(m.Subcommands))
	for i, cmd := range m.Subcommands {
		if m.Cursor == i {
			menuLines[i] = components.CursorStyle.Render("▶ " + cmd)
		} else {
			menuLines[i] = "  " + cmd
		}
	}

	menuWidth := m.Width / 4
	menu := lipgloss.JoinVertical(lipgloss.Left, menuLines...)
	menu = components.MenuStyle.Copy().Width(menuWidth).Render(menu)

	output := ""
	if m.ActiveView != nil {
		output = m.ActiveView.View()
	}
	outputWidth := m.Width - menuWidth - 3
	outputBox := components.OutputBoxStyle.Copy().Width(outputWidth).Render(output)

	body := lipgloss.JoinHorizontal(lipgloss.Top, menu, outputBox)

	footer := components.RenderFooter("↑/↓ Navigate  •  Enter Run  •  q Quit", m.Width)

	return lipgloss.JoinVertical(lipgloss.Top, header, body, footer)
}
