package pkg

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Container holds the parsed state of a single container from `container list`.
type Container struct {
	Name       string
	Status     string  // normalized: "running", "stopped", or "error"
	CPUPercent float64
	MemoryMB   float64
}

// FetchContainersMsg is the message returned by FetchContainers when the CLI
// call completes.
type FetchContainersMsg []Container

// normalizeStatus maps the raw status string from the CLI to one of the three
// canonical values: "running", "error", or "stopped".
func normalizeStatus(raw string) string {
	lower := strings.ToLower(raw)
	if strings.Contains(lower, "running") {
		return "running"
	}
	if strings.Contains(lower, "error") || strings.Contains(lower, "fail") {
		return "error"
	}
	return "stopped"
}

// FetchContainers returns a tea.Cmd that shells out to `container list`,
// parses the output into typed Container values, and returns a
// FetchContainersMsg.
//
// Parsing is defensive: malformed lines are skipped; missing CPU/memory fields
// default to 0.0.
func FetchContainers() tea.Cmd {
	return func() tea.Msg {
		out, _ := RunCommand("container", "list")

		var containers []Container

		lines := strings.Split(out, "\n")
		// First line is the header — skip it. Also skip empty lines.
		for i, line := range lines {
			if i == 0 {
				continue // header row
			}
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue // need at least NAME and STATUS
			}

			c := Container{
				Name:   fields[0],
				Status: normalizeStatus(fields[1]),
			}

			// CPU% is best-effort (field index 2)
			if len(fields) > 2 {
				if v, err := strconv.ParseFloat(fields[2], 64); err == nil {
					c.CPUPercent = v
				}
			}

			// Memory (MB) is best-effort (field index 3)
			if len(fields) > 3 {
				if v, err := strconv.ParseFloat(fields[3], 64); err == nil {
					c.MemoryMB = v
				}
			}

			containers = append(containers, c)
		}

		return FetchContainersMsg(containers)
	}
}
