package utils

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
)

type Item struct {
	title       string
	dir         string
	description string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) FilterValue() string { return i.title }
func (i Item) Dir() string         { return i.dir }

var (
	gitFolder             = ".git"
	foldersToIgnore       = []string{"node_modules"}
	enableFoldersToIgnore = false
	debugFindRepos        = false
)

func CreateItems(dir string) []list.Item {
	repoList := repoList[list.Item]{Items: []list.Item{}, createItem: func(filepath string) list.Item {
		split := strings.Split(filepath, "/")
		folderName := split[len(split)-1]

		return Item{title: folderName, dir: filepath}
	}}

	if debugFindRepos {
		measureFuncTime(func() {
			repoList.findRepos(dir)
		})
	} else {
		repoList.findRepos(dir)
	}

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
		return
	}

	for _, folder := range folders {
		isGitRepo := getIsGitRepo(folder.Name())

		if isGitRepo {
			lst.add(lst.createItem(dir))
			return
		}

		if enableFoldersToIgnore {
			for _, ignore := range foldersToIgnore {
				if folder.Name() == ignore {
					return
				}
			}
		}

		lst.findRepos(lst.getFullPath(dir, folder))
	}
}

func getIsGitRepo(folderName string) bool {
	return folderName == gitFolder
}

type CallbackFunc = func()

func measureFuncTime(x CallbackFunc) {
	start := time.Now()
	x()
	timeElapsed := time.Since(start)
	fmt.Println("Time elapsed: ", timeElapsed)
}
