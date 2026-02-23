package components

import (
	"fmt"
	"math"
	"strings"

	"container-tui/pkg"

	"github.com/charmbracelet/lipgloss"
)

// renderResourceBar renders a fixed-width progress bar using block characters.
// filled = Round((value/maxValue)*width), clamped to [0, width].
// If maxValue is 0 or value is 0, returns an all-empty bar.
func renderResourceBar(value, maxValue float64, width int) string {
	if width <= 0 {
		return ""
	}

	filled := 0
	if maxValue > 0 && value > 0 {
		filled = int(math.Round((value / maxValue) * float64(width)))
		if filled < 0 {
			filled = 0
		}
		if filled > width {
			filled = width
		}
	}

	empty := width - filled

	var sb strings.Builder
	if filled > 0 {
		sb.WriteString(ResourceBarFilledStyle.Render(strings.Repeat("█", filled)))
	}
	if empty > 0 {
		sb.WriteString(ResourceBarEmptyStyle.Render(strings.Repeat("░", empty)))
	}
	return sb.String()
}

// RenderContainerRow renders one container as a single-line row for the left
// panel. selected controls whether the row uses the highlighted style.
func RenderContainerRow(c pkg.Container, selected bool, width int) string {
	// Status dot with color.
	color := StatusColor(c.Status)
	dot := lipgloss.NewStyle().Foreground(color).Render("● ")

	// Maximum name width: width - 2 (dot) - 2 (spaces) - 8 (cpu bar) - 1 (space) - 8 (mem bar).
	maxNameWidth := width - 21
	if maxNameWidth < 1 {
		maxNameWidth = 1
	}

	name := c.Name
	if len([]rune(name)) > maxNameWidth {
		name = string([]rune(name)[:maxNameWidth-1]) + "…"
	}
	// Pad name to maxNameWidth for consistent alignment.
	name = fmt.Sprintf("%-*s", maxNameWidth, name)

	cpuBar := renderResourceBar(c.CPUPercent, 100.0, 8)
	memBar := renderResourceBar(c.MemoryMB, 1024.0, 8)

	row := dot + name + "  " + cpuBar + " " + memBar

	if selected {
		return ContainerRowSelectedStyle.Render(row)
	}
	return ContainerRowStyle.Render(row)
}
