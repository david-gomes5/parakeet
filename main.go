package main

import (
	"fmt"
	"os"
	"repo-selector/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := utils.NewModel()

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
