package system

import (
	"container-tui/ui/system/subcommands"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Cursor      int
	Subcommands []string
	ActiveView  tea.Model
	Width       int
	Height      int
}

func NewSystemModel(subs []string) *Model {
	model := &Model{
		Cursor:      0,
		Subcommands: subs,
		ActiveView:  subcommands.NewStatusModel(),
	}
	return model
}

func (m *Model) Init() tea.Cmd {
	if m.ActiveView != nil {
		return m.ActiveView.Init()
	}
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Subcommands)-1 {
				m.Cursor++
			}
		case "enter":
			switch m.Subcommands[m.Cursor] {
			case "status":
				m.ActiveView = subcommands.NewStatusModel()
			case "start":
				m.ActiveView = subcommands.NewStartModel()
			case "stop":
				m.ActiveView = subcommands.NewStopModel()
			}
			if m.ActiveView != nil {
				return m, m.ActiveView.Init()
			}
		}

	default:
		if m.ActiveView != nil {
			var cmd tea.Cmd
			m.ActiveView, cmd = m.ActiveView.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m *Model) View() string {
	return RenderSystem(m)
}
