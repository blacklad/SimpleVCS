package lib

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/util"
)

var exists = util.VCSExists(".svcs")
var currentTime = util.GetTime()
var currentUser, _ = user.Current()
var currentPath, _ = os.Getwd()
var branchesPath = path.Join(".svcs", "branches.txt")
var historyPath = path.Join(".svcs", "history")
var filesPath = path.Join(".svcs", "files")
var parentSum, _ = ioutil.ReadFile(branchesPath)
var message = "author " + currentUser.Username + "\ntime " + currentTime + "\nparent " + fmt.Sprintf("%x", parentSum)
var sum = sha1.Sum([]byte(message))
var sumString = fmt.Sprintf("%x", sum)
var branchesFile, _ = os.Create(branchesPath)
var infoFile, _ = os.Create(path.Join(historyPath, sumString+".txt"))
var filesFile, _ = os.Create(path.Join(historyPath, sumString+"_files.txt"))

func Commit() error {
	if !exists {
		return errors.New("not initialized")
	}
	branchesFile.WriteString(sumString)
	infoFile.WriteString(message)
	filepath.Walk(".", visit)
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
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	contentSum := fmt.Sprintf("%x", sha1.Sum(fileContent))
	newPath := path.Join(filesPath, contentSum)
	newFile, _ := os.Create(newPath)
	newFile.Write(fileContent)
	filesFile.WriteString(relativePath + " " + contentSum + "\n")
	return nil
}
