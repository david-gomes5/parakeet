package utils

import (
	"log"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func NewListDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.ShowDescription = false

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var (
			title string
			dir   string
		)

		if i, ok := m.SelectedItem().(Item); ok {
			title = i.title
			dir = i.dir
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.open):
				editorCmd := currentEditor.getCurrentEditor().cmd
				cmd := exec.Command(editorCmd, dir)

				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}

				return m.NewStatusMessage(StatusMessageStyle("Opening " + title + "..."))

			case key.Matches(msg, keys.toggle):
				currentEditor.incrementCurrentEditorIndex()
			}
		}

		return nil
	}

	help := []key.Binding{keys.open, keys.toggle}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}
