package utils

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ItemStyle          = lipgloss.NewStyle().PaddingLeft(2)
	ViewStyle          = lipgloss.NewStyle().Margin(5)
	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Bold(true).
				Render
)
