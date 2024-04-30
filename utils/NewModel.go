package utils

import (
	"github.com/charmbracelet/bubbles/key"
	bubbleList "github.com/charmbracelet/bubbles/list"
)

func NewModel() Model {
	var (
		delegateKeys = NewDelegateKeyMap()
		items        = CreateItems("/Users/davidgomes/repos")
		listKeys     = NewListKeyMap()
	)

	// Setup list
	delegate := NewItemDelegate(delegateKeys)
	list := bubbleList.New(items, delegate, 0, 0)
	list.Title = "Repositories"
	list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.InsertItem,
			listKeys.ToggleSpinner,
		}
	}

	return Model{List: list}
}
