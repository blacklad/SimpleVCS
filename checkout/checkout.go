package checkout

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MSathieu/SimpleVCS/util"
)

var exists = util.VCSExists(".svcs")
var filesPath = ".svcs/files"
var historyPath = ".svcs/history"
var currentPath, _ = os.Getwd()
var branchesPath = ".svcs/branches.txt"

func Checkout(commitHash string) error {
	if !exists {
		return errors.New("not initialized")
	}
	filesEntryPath := path.Join(historyPath, commitHash+"_files.txt")
	filesContent, _ := ioutil.ReadFile(filesEntryPath)
	files := strings.Split(string(filesContent), "\n")
	for _, fileEntry := range files {
		if fileEntry == "" {
			break
		}
		mapping := strings.Split(fileEntry, " ")
		copyFrom := path.Join(filesPath, mapping[1])
		fileContent, _ := ioutil.ReadFile(copyFrom)
		splitFileArr := strings.Split(mapping[0], "/")
		splitFileArr = splitFileArr[:len(splitFileArr)-1]
		toDir := ""
		for _, element := range splitFileArr {
			toDir = toDir + element + "/"
		}
		os.MkdirAll(toDir, 666)
		newFile, _ := os.Create(mapping[0])
		newFile.Write(fileContent)
	}
	return nil
}
