package lib

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/util"
)

var currentTime = util.GetTime()
var currentUser, _ = user.Current()
var currentPath, _ = os.Getwd()
var branchesPath = path.Join(".svcs", "branches.txt")
var currentSum, _ = ioutil.ReadFile(branchesPath)
var message = "author " + currentUser.Username + "\ntime " + currentTime + "\nparent " + fmt.Sprintf("%x", currentSum)
var sum = sha1.Sum([]byte(message))
var commitPath = path.Join(".svcs", fmt.Sprintf("%x", sum))
var messagePath = path.Join(commitPath, "commit-info.txt")

func Commit() error {
	exists := util.VCSExists(".svcs")
	if !exists {
		return errors.New("not initialized")
	}
	branchesFile, _ := os.Create(branchesPath)
	branchesFile.Write([]byte(fmt.Sprintf("%x", sum)))
	os.Mkdir(commitPath, 0700)
	file, _ := os.Create(messagePath)
	file.Write([]byte(message))
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
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	newPath := path.Join(commitPath, relativePath)
	if fileInfo.IsDir() {
		os.Mkdir(newPath, fileInfo.Mode())
	} else {
		newFile, _ := os.Create(newPath)
		oldFile, _ := os.Open(fixedPath)
		io.Copy(newFile, oldFile)
	}
	return nil
}
