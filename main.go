package main

import (
	"flag"
	"fmt"
	"os"
	"parakeet/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	folderPaths := constructFlags()

	if len(folderPaths) <= 0 {
		fmt.Println("Unable to run program without a folder path")
		os.Exit(1)
	}

	m := utils.NewModel(folderPaths)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func constructFlags() utils.FolderPaths {
	var folderPaths utils.FolderPaths

	flag.Var(&folderPaths, "f", "Folder paths to open")
	flag.Parse()

	return folderPaths
}
