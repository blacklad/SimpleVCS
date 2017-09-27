package cmd

import (
	"os"
	"path/filepath"

	"github.com/MSathieu/SimpleVCS/lib"
)

var currentFiles []string

//Status prints the status.
func Status() error {
	head, err := lib.GetHead()
	if err != nil {
		return err
	}
	commit, err := lib.GetCommit(head)
	if err != nil {
		return err
	}
	files, err := commit.GetFiles()
	if err != nil {
		return err
	}
	err = filepath.Walk(".", statusVisit)
	if err != nil {
		return err
	}
	return nil
}
func statusVisit(filePath string, fileInfo os.FileInfo, err error) error {
	return nil
}
