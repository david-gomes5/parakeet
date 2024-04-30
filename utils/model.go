package utils

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	List list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.List.Update(msg)
	m.List = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return appStyle.Render(m.List.View())
}
