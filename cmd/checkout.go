package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Checkout(commitHash string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	filesEntryPath := path.Join(".svcs/history", commitHash+"_files.txt")
	filesContent, _ := ioutil.ReadFile(filesEntryPath)
	files := strings.Split(string(filesContent), "\n")
	for _, fileEntry := range files {
		if fileEntry == "" {
			continue
		}
		mapping := strings.Split(fileEntry, " ")
		copyFrom := path.Join(".svcs/files", mapping[1])
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
