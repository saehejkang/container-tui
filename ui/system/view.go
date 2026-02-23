package system

import (
	"container-tui/ui/components"

	"github.com/charmbracelet/lipgloss"
)

func RenderSystem(m *Model) string {
	header := components.RenderHeaderWithStatus("Container TUI", false)

	body := ""
	if m.Width > 0 && m.Height > 0 {
		if m.Loading {
			body = "Loading containers..."
		} else if m.Error != "" {
			body = "Error: " + m.Error
		} else {
			body = renderContainerList(m)
		}
	}

	footer := components.RenderFooter("↑/↓ Navigate  •  ? Help  •  q Quit", m.Width)

	return lipgloss.JoinVertical(lipgloss.Top, header, body, footer)
}

func renderContainerList(m *Model) string {
	if len(m.Containers) == 0 {
		return "No containers found."
	}

	lines := make([]string, len(m.Containers))
	for i, c := range m.Containers {
		prefix := "  "
		if i == m.SelectedIndex {
			prefix = components.CursorStyle.Render("▶ ")
		}
		lines[i] = prefix + c.Name + "  [" + c.Status + "]"
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
