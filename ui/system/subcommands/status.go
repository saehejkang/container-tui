package subcommands

import (
	"fmt"
	"strings"

	"container-tui/pkg"

	tea "github.com/charmbracelet/bubbletea"
)

type StatusModel struct {
	Fields map[string]string
	Output []string
}

type StatusMsg map[string]string

func NewStatusModel() *StatusModel {
	fields := map[string]string{
		"status":            "",
		"appRoot":           "",
		"installRoot":       "",
		"logRoot":           "",
		"apiserver.version": "",
		"apiserver.commit":  "",
		"apiserver.build":   "",
		"apiserver.appName": "",
	}

	return &StatusModel{
		Fields: fields,
		Output: []string{"Fetching container system status..."},
	}
}

func (m *StatusModel) Init() tea.Cmd {
	return func() tea.Msg {
		out, _ := pkg.RunCommand("container", "system", "status")
		lines := strings.Split(out, "\n")
		for _, line := range lines {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			key := parts[0]
			value := strings.Join(parts[1:], " ")
			m.Fields[key] = value
		}
		return StatusMsg(m.Fields)
	}
}

func (m *StatusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if status, ok := msg.(StatusMsg); ok {
		m.Output = []string{}
		for _, key := range []string{
			"status", "appRoot", "installRoot", "logRoot",
			"apiserver.version", "apiserver.commit", "apiserver.build", "apiserver.appName",
		} {
			val := status[key]
			if val == "" {
				val = "<not running>"
			}
			m.Output = append(m.Output, fmt.Sprintf("%-20s %s", key, val))
		}
	}
	return m, nil
}

func (m *StatusModel) View() string {
	return strings.Join(m.Output, "\n")
}
