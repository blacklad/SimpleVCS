package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Checkout checks out the specified commit.
func Checkout(commitHash string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	filesEntryPath := path.Join(".svcs/history", commitHash+"_files.txt")
	filesContent, err := ioutil.ReadFile(filesEntryPath)
	if err != nil {
		return err
	}
	files := strings.Split(string(filesContent), "\n")
	for _, fileEntry := range files {
		if fileEntry == "" {
			continue
		}
		mapping := strings.Split(fileEntry, " ")
		copyFrom := path.Join(".svcs/files", mapping[1])
		fileContent, err := ioutil.ReadFile(copyFrom)
		if err != nil {
			return err
		}
		splitFileArr := strings.Split(mapping[0], "/")
		splitFileArr = splitFileArr[:len(splitFileArr)-1]
		toDir := ""
		for _, element := range splitFileArr {
			toDir = toDir + element + "/"
		}
		err = os.MkdirAll(toDir, 666)
		if err != nil {
			return err
		}
		newFile, err := os.Create(mapping[0])
		if err != nil {
			return err
		}
		unzippedContent := lib.Unzip(fileContent)
		_, err = newFile.WriteString(unzippedContent)
		if err != nil {
			return err
		}
	}
	return nil
}
