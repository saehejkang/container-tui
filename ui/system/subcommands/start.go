package subcommands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type StartModel struct {
	progress float64
	done     bool
	err      error
}

type startFinishedMsg struct {
	err error
}
type tickMsg struct{}

func NewStartModel() *StartModel {
	return &StartModel{}
}

func (m *StartModel) Init() tea.Cmd {
	return tea.Batch(runStartCommand(), tickProgress())
}

func (m *StartModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if !m.done && m.progress < 0.95 {
			m.progress += 0.02
		}
		return m, tickProgress()

	case startFinishedMsg:
		m.err = msg.err
		m.progress = 1.0
		m.done = true
		return m, nil

	default:
		return m, nil
	}
}

func (m *StartModel) View() string {
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
			return "❌ Container Start Failed\n" + bar
		}
		return "✅ Container Started!\n" + bar
	}

	return "Starting Container...\n" + bar
}

func tickProgress() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(time.Time) tea.Msg { return tickMsg{} })
}

func runStartCommand() tea.Cmd {
	return func() tea.Msg {
		err := exec.Command("container", "system", "start").Run()
		return startFinishedMsg{err: err}
	}
}
