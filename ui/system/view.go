package system

import (
	"container-tui/ui/components"

	"github.com/charmbracelet/lipgloss"
)

func RenderSystem(m *Model) string {
	if m.Width == 0 || m.Height == 0 {
		return "Initializing..."
	}

	leftWidth := m.Width / 3
	rightWidth := m.Width - leftWidth - 2 // -2 for border/margin

	// Derive isRunning from container statuses.
	isRunning := false
	for _, c := range m.Containers {
		if c.Status == "running" {
			isRunning = true
			break
		}
	}
	header := components.RenderHeaderWithStatus("Container TUI", isRunning)

	// Build left panel (container list).
	var leftContent string
	if m.Loading {
		leftContent = "  Loading..."
	} else if len(m.Containers) == 0 {
		leftContent = "  No containers"
	} else {
		rows := make([]string, len(m.Containers))
		for i, c := range m.Containers {
			rows[i] = components.RenderContainerRow(c, i == m.SelectedIndex, leftWidth)
		}
		leftContent = lipgloss.JoinVertical(lipgloss.Left, rows...)
	}
	leftPanel := components.MenuStyle.Copy().Width(leftWidth).Render(leftContent)

	// Build right panel placeholder.
	placeholder := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render("Select a container to view details")
	rightPanel := components.OutputBoxStyle.Copy().Width(rightWidth).Render(placeholder)

	// Compose body.
	body := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

	// Render footer.
	footer := components.RenderFooter("↑/↓  Navigate  •  ?  Help  •  q  Quit", m.Width)

	// Base layout.
	baseLayout := lipgloss.JoinVertical(lipgloss.Top, header, body, footer)

	// Overlay help if needed.
	if m.ShowHelp {
		overlay := components.RenderHelpOverlay(m.Width, m.Height)
		return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, overlay)
	}

	return baseLayout
}
