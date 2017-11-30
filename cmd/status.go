package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

var currentFiles []lib.TreeFile

//Status prints the status.
func Status() error {
	head, err := lib.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("can't view status in detached state")
	}
	fmt.Println("branch " + head.Branch.Name)
	commit := head.Branch.Commit
	err = filepath.Walk(".", statusVisit)
	if err != nil {
		return err
	}
	changes := lib.GenerateChange(commit.Tree.Files, currentFiles)
	for _, change := range changes {
		fmt.Println(change.Type + " " + change.Name)
	}
	return nil
}
func statusVisit(filePath string, fileInfo os.FileInfo, err error) error {
	fixedPath := filepath.ToSlash(filePath)
	pathArr := strings.Split(fixedPath, "/")
	for _, pathPart := range pathArr {
		ignored, err := lib.CheckIgnored(pathPart)
		if err != nil {
			return err
		}
		if ignored {
			return nil
		}
	}
	if fileInfo.IsDir() {
		return nil
	}
	currentFileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	checksum := gotils.GetChecksum(string(currentFileContent))
	currentFiles = append(currentFiles, lib.TreeFile{Name: fixedPath, File: vcsfile.File{Hash: checksum}})
	return nil
}
