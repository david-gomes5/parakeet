package utils

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ViewStyle  = lipgloss.NewStyle().Margin(2)
	TitleStyle = lipgloss.NewStyle().Margin(0).Padding(0, 1).Foreground(lipgloss.Color("#FF00FF")).Border(lipgloss.RoundedBorder())
	GetTitle   = func() string {
		return lipgloss.JoinVertical(lipgloss.Left, "Repositories",
			currentEditor.getCurrentEditor().name,
		)
	}
	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Bold(true).
				Render
)
