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
)

var currentFiles []string

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
	files := commit.GetFiles()
	err = filepath.Walk(".", statusVisit)
	if err != nil {
		return err
	}
	var changes []string
	for _, line := range files {
		change := "deleted"
		changed := true
		split := strings.Split(line, " ")
		for _, currentLine := range currentFiles {
			currentSplit := strings.Split(currentLine, " ")
			if split[0] == currentSplit[0] {
				if split[1] == currentSplit[1] {
					changed = false
				} else {
					change = "changed"
				}
			}
		}
		if changed {
			changes = append(changes, change+" "+split[0])
		}
	}
	for _, currentLine := range currentFiles {
		change := "created"
		changed := true
		currentSplit := strings.Split(currentLine, " ")
		for _, line := range files {
			split := strings.Split(line, " ")
			if split[0] == currentSplit[0] {
				changed = false
			}
		}
		if changed {
			changes = append(changes, change+" "+currentSplit[0])
		}
	}
	fmt.Println(strings.Join(changes, "\n"))
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
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	currentFiles = append(currentFiles, relativePath+" "+checksum)
	return nil
}
