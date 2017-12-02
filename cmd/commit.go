package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/SimpleVCS/ignore"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

type safeFilesSlice struct {
	files []string
	mutex sync.Mutex
}

var filesStruct safeFilesSlice
var commitWait sync.WaitGroup

//Commit commits the current directory.
func Commit(message string) error {
	if message == "" {
		return errors.New("you must specify a message")
	}
	head, err := util.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("cannot commit in detached state")
	}
	err = util.ExecHook("precommit")
	if err != nil {
		return err
	}
	err = filepath.Walk(".", commitVisit)
	if err != nil {
		return err
	}
	commitWait.Wait()
	sumString, err := vcscommit.Create(message, filesStruct.files)
	if err != nil {
		return err
	}
	err = vcsbranch.Update(head.Branch, sumString)
	if err != nil {
		return err
	}
	err = util.ExecHook("postcommit")
	if err != nil {
		return err
	}
	return nil
}
func commitVisit(filePath string, fileInfo os.FileInfo, err error) error {
	commitWait.Add(1)
	go concCommitVisit(filePath, fileInfo)
	return nil
}
func concCommitVisit(filePath string, fileInfo os.FileInfo) {
	defer commitWait.Done()
	fixedPath := filepath.ToSlash(filePath)
	pathArr := strings.Split(fixedPath, "/")
	for _, pathPart := range pathArr {
		ignored, err := ignore.CheckIgnored(pathPart)
		if err != nil {
			log.Fatal(err)
		}
		if ignored {
			return
		}
	}
	if fileInfo.IsDir() {
		return
	}
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileObj := vcsfile.File{Content: string(fileContent)}
	err = fileObj.Save()
	if err != nil {
		log.Fatal(err)
	}
	filesStruct.mutex.Lock()
	defer filesStruct.mutex.Unlock()
	filesStruct.files = append(filesStruct.files, fixedPath+" "+fileObj.Hash)
}
