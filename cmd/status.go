package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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
	checksum := lib.GetChecksum(string(currentFileContent))
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	currentFiles = append(currentFiles, relativePath+" "+checksum)
	return nil
}
