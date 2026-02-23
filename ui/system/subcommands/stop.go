package subcommands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type StopModel struct {
	progress float64
	done     bool
	err      error
}

type stopFinishedMsg struct {
	err error
}
type tickStopMsg struct{}

func NewStopModel() *StopModel {
	return &StopModel{}
}

func (m *StopModel) Init() tea.Cmd {
	return tea.Batch(runStopCommand(), tickProgressStop())
}

func (m *StopModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickStopMsg:
		if !m.done && m.progress < 0.95 {
			m.progress += 0.02
		}
		return m, tickProgressStop()

	case stopFinishedMsg:
		m.err = msg.err
		m.progress = 1.0
		m.done = true
		return m, nil

	default:
		return m, nil
	}
}

func (m *StopModel) View() string {
	barWidth := 40
	filled := int(m.progress * float64(barWidth))
	empty := barWidth - filled

	bar := fmt.Sprintf("[%s%s] %d%%",
		strings.Repeat("█", filled),
		strings.Repeat(" ", empty),
		int(m.progress*100),
	)

	if m.done {
		if m.err != nil {
			return "❌ Container Stop Failed\n" + bar
		}
		return "🛑 Container Stopped!\n" + bar
	}

	return "Stopping Container...\n" + bar
}

func tickProgressStop() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(time.Time) tea.Msg { return tickStopMsg{} })
}

func runStopCommand() tea.Cmd {
	return func() tea.Msg {
		err := exec.Command("container", "system", "stop").Run()
		return stopFinishedMsg{err: err}
	}
}
