package utils

import "fmt"

type FolderPaths []string

func (i *FolderPaths) String() string {
	return fmt.Sprint(*i)
}

func (i *FolderPaths) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *FolderPaths) FolderPaths() []string { return *i }
