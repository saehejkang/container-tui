package ui

import tea "github.com/charmbracelet/bubbletea"

type TUI struct {
	CurrentView tea.Model
}

func (t *TUI) Init() tea.Cmd {
	if t.CurrentView != nil {
		return t.CurrentView.Init()
	}
	return nil
}

func (t *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if t.CurrentView != nil {
		updated, cmd := t.CurrentView.Update(msg)
		t.CurrentView = updated
		return t, cmd
	}
	return t, nil
}

func (t *TUI) View() string {
	if t.CurrentView != nil {
		return t.CurrentView.View()
	}
	return ""
}
