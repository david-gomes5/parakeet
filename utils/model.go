package utils

import (
	bubbleList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	List bubbleList.Model
}

func (m Model) Init() tea.Cmd {
	// TODO: implement reading from a config file?
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := DocStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return DocStyle.Render(m.List.View())
}

func NewModel() Model {
	var (
		delegateKeys = NewDelegateKeyMap()
		items        = CreateItems("/Users/davidgomes/repos")
	)

	// Setup list
	delegate := NewListDelegate(delegateKeys)
	list := bubbleList.New(items, delegate, 0, 0)
	list.Title = "Repositories"

	return Model{List: list}
}
