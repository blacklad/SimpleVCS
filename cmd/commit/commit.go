package commit

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/SimpleVCS/cmd/ignore"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//SafeFilesSlice is the files slice object
type SafeFilesSlice struct {
	Files []string
	mutex sync.Mutex
}

//FilesStruct is the filesstruct var
var FilesStruct SafeFilesSlice

//CommitWait is the commit waitgroup
var CommitWait sync.WaitGroup

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
	err = filepath.Walk(".", CommitVisit)
	if err != nil {
		return err
	}
	CommitWait.Wait()
	sumString, err := types.CreateCommit(message, FilesStruct.Files)
	if err != nil {
		return err
	}
	err = types.UpdateBranch(head.Branch, sumString)
	if err != nil {
		return err
	}
	err = util.ExecHook("postcommit")
	if err != nil {
		return err
	}
	return nil
}

//CommitVisit visits the directory and creates the files
func CommitVisit(filePath string, fileInfo os.FileInfo, err error) error {
	CommitWait.Add(1)
	go concCommitVisit(filePath, fileInfo)
	return nil
}
func concCommitVisit(filePath string, fileInfo os.FileInfo) {
	defer CommitWait.Done()
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
	fileObj := types.File{Content: string(fileContent)}
	err = fileObj.Save()
	if err != nil {
		log.Fatal(err)
	}
	FilesStruct.mutex.Lock()
	defer FilesStruct.mutex.Unlock()
	FilesStruct.Files = append(FilesStruct.Files, fixedPath+" "+fileObj.Hash)
}
