package main

import (
	"fmt"
	"os"

	"container-tui/ui/system"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	systemModel := system.NewSystemModel()

	p := tea.NewProgram(systemModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		os.Exit(1)
	}
}
