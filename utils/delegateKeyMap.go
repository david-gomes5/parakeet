package utils

import "github.com/charmbracelet/bubbles/key"

type delegateKeyMap struct {
	choose key.Binding
	toggle key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.toggle,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.toggle,
		},
	}
}

func NewDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		toggle: key.NewBinding(
			key.WithKeys("x"),
			key.WithHelp("x", "switch IDE"),
		),
	}
}
