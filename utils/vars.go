package utils

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	// TitleStyle         = lipgloss.NewStyle().MarginLeft(2).
	ItemStyle          = lipgloss.NewStyle().PaddingLeft(2)
	DocStyle           = lipgloss.NewStyle().Margin(1, 2)
	SelectedItemStyle  = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle    = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle          = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle      = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
	appStyle = lipgloss.NewStyle().Padding(1, 2)
)
