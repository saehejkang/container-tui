package system

import (
	"time"

	"container-tui/pkg"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Containers    []pkg.Container
	SelectedIndex int
	ShowHelp      bool
	Width         int
	Height        int
	Loading       bool
	Error         string
}

func NewSystemModel() *Model {
	return &Model{
		Loading: true,
	}
}

type refreshTickMsg struct{}

func tickRefresh() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return refreshTickMsg{}
	})
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		pkg.FetchContainers(),
		tickRefresh(),
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width > 0 && msg.Height > 0 {
			m.Width = msg.Width
			m.Height = msg.Height
		}
		return m, nil

	case pkg.FetchContainersMsg:
		m.Containers = []pkg.Container(msg)
		m.Loading = false
		if len(m.Containers) == 0 {
			m.SelectedIndex = 0
		} else if m.SelectedIndex >= len(m.Containers) {
			m.SelectedIndex = len(m.Containers) - 1
		}
		return m, nil

	case refreshTickMsg:
		return m, tea.Batch(pkg.FetchContainers(), tickRefresh())

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "?":
			m.ShowHelp = !m.ShowHelp
		case "esc":
			if m.ShowHelp {
				m.ShowHelp = false
			}
		case "up", "k":
			if !m.ShowHelp && m.SelectedIndex > 0 {
				m.SelectedIndex--
			}
		case "down", "j":
			if !m.ShowHelp && m.SelectedIndex < len(m.Containers)-1 {
				m.SelectedIndex++
			}
		}
		return m, nil

	default:
		return m, nil
	}
}

func (m *Model) View() string {
	return RenderSystem(m)
}
