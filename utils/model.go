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
	// TODO: Check if IDE commands exist, if not, remove them from the list
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := ViewStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return ViewStyle.Render(m.List.View())
}

func NewModel(folderPaths FolderPaths) Model {
	var (
		delegateKeys = NewDelegateKeyMap()
		items        = CreateItems(folderPaths.FolderPaths())
	)

	// Setup list
	delegate := NewListDelegate(delegateKeys)
	list := bubbleList.New(items, delegate, 0, 0)

	// Set styles
	list.Styles.Title = TitleStyle

	// Set title
	title := GetTitle()
	list.Title = title

	return Model{List: list}
}
