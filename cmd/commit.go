package cmd

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

var files []string

//Commit commits the current directory.
func Commit(message string) error {
	head, err := ioutil.ReadFile(".svcs/head.txt")
	if err != nil {
		return err
	}
	if string(head) == "DETACHED" {
		return errors.New("cannot commit in detached state")
	}
	branch := string(head)
	err = filepath.Walk(".", visit)
	if err != nil {
		return err
	}
	sumString, err := lib.Commit(message, files)
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
		if pathPart == ".svcs" || pathPart == ".git" || pathPart == ".hg" || pathPart == ".svn" {
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
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	contentSum := fmt.Sprintf("%x", sha1.Sum(fileContent))
	newPath := path.Join(".svcs/files", contentSum)
	zippedContent := lib.Zip(fileContent)
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	_, err = newFile.WriteString(zippedContent)
	if err != nil {
		return err
	}
	files = append(files, relativePath+" "+contentSum)
	return err
}
