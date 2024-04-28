package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type item struct {
	title       string
	description string
}

func (i item) FilterValue() string { return i.title }

func isGitRepo(folders []fs.DirEntry) bool {
	for _, folder := range folders {
		if folder.Name() == ".git" {
			return true
		}
	}

	return false
}

// // TODO: handle recursive directories when git not found
func createItems(dir string) []list.Item {
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
				list = append(list, item{title: folder.Name()})
			}
		}
	}

	return list
}
