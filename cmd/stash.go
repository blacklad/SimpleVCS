package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/util"
)

var wait sync.WaitGroup

//CreateStash creates a stash
func CreateStash(name string) error {
	err := filepath.Walk(".", commitVisit)
	if err != nil {
		return err
	}
	commitWait.Wait()
	stashFile, err := os.Create(".svcs/stashes/" + name)
	if err != nil {
		return err
	}
	defer stashFile.Close()
	stashFileContent := strings.Join(filesStruct.files, "\n")
	stashFileContent, err = util.Zip(stashFileContent)
	if err != nil {
		return err
	}
	_, err = stashFile.WriteString(stashFileContent)
	return err
}

//CheckoutStash checkouts a stash
func CheckoutStash(name string) error {
	stash, err := ioutil.ReadFile(".svcs/stashes/" + name)
	if err != nil {
		return err
	}
	stashContent, err := util.Unzip(string(stash))
	if err != nil {
		return err
	}
	stashArr := strings.Split(stashContent, "\n")
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
	copyFrom := path.Join(".svcs/files", hash)
	fileContent, err := ioutil.ReadFile(copyFrom)
	if err != nil {
		log.Fatal(err)
	}
	unzippedContent, err := util.Unzip(string(fileContent))
	if err != nil {
		log.Fatal(err)
	}
	err = gotils.CheckIntegrity(unzippedContent, hash)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(filepath.Dir(name), 666)
	newFile, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	newFile.WriteString(unzippedContent)
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
	return os.Remove(".svcs/stashes/" + name)
}
