package cmd

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

var currentTime = lib.GetTime()
var branch string

//Commit commits the current directory.
func Commit(currentBranch string, message string) error {
	branch = currentBranch
	info, sumString, err := lib.CreateCommitInfo(currentTime, branch)
	if err != nil {
		return err
	}
	err = lib.CreateCommit(info, sumString, message)
	if err != nil {
		return err
	}
	err = filepath.Walk(".", visit)
	if err != nil {
		return err
	}
	err = lib.UpdateBranch(currentBranch, sumString)
	return err
}
func visit(filePath string, fileInfo os.FileInfo, err error) error {
	fixedPath := strings.Replace(filePath, "\\", "/", -1)
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
	_, sumString, err := lib.CreateCommitInfo(currentTime, branch)
	if err != nil {
		return err
	}
	fileEntriesPath := path.Join(".svcs/history", sumString+"_files.txt")
	fileEntriesFile, err := os.OpenFile(fileEntriesPath, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = fileEntriesFile.WriteString(relativePath + " " + contentSum + "\n")
	return err
}
