package cmd

import (
	"bytes"
	"compress/gzip"
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

var currentTime = lib.GetTime()
var branch string

func Commit(currentBranch string, message string) error {
	branch = currentBranch
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	info, sumString := lib.CreateCommitInfo(currentTime, branch)
	lib.CreateCommit(info, sumString, message)
	filepath.Walk(".", visit)
	lib.UpdateBranch(currentBranch, sumString)
	return nil
}
func visit(filePath string, fileInfo os.FileInfo, err error) error {
	fixedPath := strings.Replace(filePath, "\\", "/", -1)
	pathArr := strings.Split(fixedPath, "/")
	for _, pathPart := range pathArr {
		if pathPart == ".svcs" {
			return nil
		}
	}
	if fileInfo.IsDir() {
		return nil
	}
	fileContent, _ := ioutil.ReadFile(filePath)
	currentPath, _ := os.Getwd()
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	contentSum := fmt.Sprintf("%x", sha1.Sum(fileContent))
	newPath := path.Join(".svcs/files", contentSum)
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	comp.Write(fileContent)
	comp.Close()
	newFile, _ := os.Create(newPath)
	compBytes.WriteTo(newFile)
	_, sumString := lib.CreateCommitInfo(currentTime, branch)
	fileEntriesPath := path.Join(".svcs/history", sumString+"_files.txt")
	fileEntriesFile, _ := os.OpenFile(fileEntriesPath, os.O_APPEND, 0666)
	fileEntriesFile.WriteString(relativePath + " " + contentSum + "\n")
	return nil
}
