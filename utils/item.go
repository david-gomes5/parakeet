package utils

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type Item struct {
	title       string
	description string
}

func CreateItems(dir string) []list.Item {
	list := []list.Item{}

	folders, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	for _, folder := range folders {
		if folder.IsDir() {
			folders, _ = os.ReadDir(dir + "/" + folder.Name())

			var isGitRepo = isGitRepo(folders)

			if isGitRepo {
				list = append(list, Item{title: folder.Name(), description: "This is a git repository"})
			}
		}
	}

	return list
}

func isGitRepo(folders []fs.DirEntry) bool {
	for _, folder := range folders {
		if folder.Name() == ".git" {
			return true
		}
	}

	return false
}

func (i Item) FilterValue() string { return i.title }
