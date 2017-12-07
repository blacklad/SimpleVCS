package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/ignore"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcschange"
)

var currentFiles []types.TreeFile

//Status prints the status.
func Status() error {
	head, err := util.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("can't view status in detached state")
	}
	fmt.Println("branch " + head.Branch)
	headBranch, err := types.GetBranch(head.Branch)
	if err != nil {
		return err
	}
	err = filepath.Walk(".", statusVisit)
	if err != nil {
		return err
	}
	changes := vcschange.GenerateChange(headBranch.Commit.Tree.Files, currentFiles)
	for _, change := range changes {
		fmt.Println(change.Type + " " + change.Name)
	}
	return nil
}
func statusVisit(filePath string, fileInfo os.FileInfo, err error) error {
	fixedPath := filepath.ToSlash(filePath)
	pathArr := strings.Split(fixedPath, "/")
	for _, pathPart := range pathArr {
		ignored, err := ignore.CheckIgnored(pathPart)
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
	currentFileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	checksum := gotils.GetChecksum(string(currentFileContent))
	currentFiles = append(currentFiles, types.TreeFile{Name: fixedPath, File: types.File{Hash: checksum}})
	return nil
}
