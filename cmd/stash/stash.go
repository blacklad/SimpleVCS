package stash

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/cmd/commit"
	"github.com/MSathieu/SimpleVCS/util"
)

var wait sync.WaitGroup

//CreateStash creates a stash
func CreateStash(name string) error {
	err := filepath.Walk(".", commit.CommitVisit)
	if err != nil {
		return err
	}
	commit.CommitWait.Wait()
	stashFileContent := strings.Join(commit.FilesStruct.Files, "\n")
	util.DB.Create(&util.Stash{Name: name, Files: stashFileContent})
	return nil
}

//CheckoutStash checkouts a stash
func CheckoutStash(name string) error {
	var stash util.Stash
	util.DB.Where(&util.Stash{Name: name}).First(&stash)
	stashArr := strings.Split(stash.Files, "\n")
	for _, fileEntry := range stashArr {
		if fileEntry == "" {
			continue
		}
		wait.Add(1)
		mapping := strings.Split(fileEntry, " ")
		go concProcessFile(mapping[1], mapping[0])
	}
	wait.Wait()
	return nil
}
func concProcessFile(hash string, name string) {
	var file util.File
	util.DB.Where(&util.File{Hash: hash}).First(&file)
	err := gotils.CheckIntegrity(file.Content, hash)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(filepath.Dir(name), 666)
	newFile, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	newFile.WriteString(file.Content)
	if err != nil {
		log.Fatal(err)
	}
	err = newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	wait.Done()
}

//RemoveStash removes a stash
func RemoveStash(name string) error {
	util.DB.Delete(&util.Stash{Name: name})
	return nil
}
