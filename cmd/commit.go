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

	"github.com/MSathieu/SimpleVCS/util"
)

var currentTime = util.GetTime()

func Commit() error {
	if !util.VCSExists(".svcs") {
		return errors.New("not initialized")
	}

	message, sumString := util.CreateMessage(currentTime)
	fileEntriesPath := path.Join(".svcs/history", sumString+"_files.txt")
	os.Create(fileEntriesPath)
	infoFile, _ := os.Create(path.Join(".svcs/history", sumString+".txt"))
	infoFile.WriteString(message)
	filepath.Walk(".", visit)
	branchesFile, _ := os.Create(".svcs/branches.txt")
	branchesFile.WriteString(sumString)
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
	newFile, _ := os.Create(newPath)
	newFile.Write(fileContent)
	_, sumString := util.CreateMessage(currentTime)
	fileEntriesPath := path.Join(".svcs/history", sumString+"_files.txt")
	fileEntriesFile, _ := os.OpenFile(fileEntriesPath, os.O_APPEND, 0666)
	fileEntriesFile.WriteString(relativePath + " " + contentSum + "\n")
	return nil
}
