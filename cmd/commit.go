package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/SimpleVCS/lib"
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
	head, err := lib.GetHead()
	if err != nil {
		return err
	}
	if head == "DETACHED" {
		return errors.New("cannot commit in detached state")
	}
	branch := head
	err = filepath.Walk(".", commitVisit)
	if err != nil {
		return err
	}
	commitWait.Wait()
	sumString, err := lib.CreateCommit(message, filesStruct.files)
	if err != nil {
		return err
	}
	err = lib.UpdateBranch(branch, sumString)
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
		ignored, err := lib.CheckIgnored(pathPart)
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
	file, err := lib.AddFile(string(fileContent))
	if err != nil {
		log.Fatal(err)
	}
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	filesStruct.mutex.Lock()
	defer filesStruct.mutex.Unlock()
	filesStruct.files = append(filesStruct.files, relativePath+" "+file.Hash)
}
