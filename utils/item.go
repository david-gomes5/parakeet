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

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) FilterValue() string { return i.title }

var gitFolder = ".git"
var foldersToIgnore = []string{"node_modules"}

func CreateItems(dir string) []list.Item {
	repoList := repoList[list.Item]{Items: []list.Item{}, createItem: func(filepath string) list.Item {
		return Item{title: filepath}
	}}

	repoList.findRepos(dir)

	return repoList.Items
}

type repoList[T any] struct {
	Items      []T
	createItem func(filepath string) T
}

func (lst *repoList[T]) add(item T) {
	lst.Items = append(lst.Items, item)
}

func (lst *repoList[T]) getFullPath(dir string, folder fs.DirEntry) string {
	return dir + "/" + folder.Name()
}

func (lst *repoList[T]) findRepos(dir string) {
	folders, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	for _, folder := range folders {
		isGitRepo := getIsGitRepo(folder.Name())

		if isGitRepo {
			lst.add(lst.createItem(dir + "/" + folder.Name()))
			return
		}
	}

	for _, folder := range folders {
		lst.findRepos(lst.getFullPath(dir, folder))
	}
}

func getIsGitRepo(folderName string) bool {
	return folderName == gitFolder
}
