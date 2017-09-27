package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

var files []string

//Commit commits the current directory.
func Commit(message string) error {
	if message == "" {
		return errors.New("you must specify a message")
	}
	head, err := lib.GetHead()
	if err != nil {
		return err
	}
	if head == "DETACHED" {
		return errors.New("cannot commit in detached state")
	}
	branch := head
	err = filepath.Walk(".", visit)
	if err != nil {
		return err
	}
	sumString, err := lib.CreateCommit(message, files)
	if err != nil {
		return err
	}
	err = lib.UpdateBranch(branch, sumString)
	return err
}
func visit(filePath string, fileInfo os.FileInfo, err error) error {
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
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	file, err := lib.AddFile(string(fileContent))
	if err != nil {
		return err
	}
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	files = append(files, relativePath+" "+file.Hash)
	return err
}
