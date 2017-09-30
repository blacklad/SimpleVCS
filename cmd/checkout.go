package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/MSathieu/SimpleVCS/lib"
)

var wait sync.WaitGroup

//Checkout checks out the specified commit.
func Checkout(commitHash string) error {
	checkoutBranch := commitHash
	currentCommit, isBranch, err := lib.ConvertToCommit(commitHash)
	if err != nil {
		return err
	}
	commitHash = currentCommit.Hash
	commit, err := lib.GetCommit(commitHash)
	if err != nil {
		return err
	}
	files, err := commit.GetFiles()
	if err != nil {
		return err
	}
	for _, fileEntry := range files {
		if fileEntry == "" {
			continue
		}
		wait.Add(1)
		mapping := strings.Split(fileEntry, " ")
		go concProcessFile(mapping[1], mapping[0])
	}
	head, err := os.Create(".svcs/head.txt")
	if err != nil {
		return err
	}
	if isBranch {
		_, err = head.WriteString(checkoutBranch)
	} else {
		_, err = head.WriteString("DETACHED")
	}
	if err != nil {
		return err
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
	splitFileArr := strings.Split(name, "/")
	splitFileArr = splitFileArr[:len(splitFileArr)-1]
	toDir := ""
	for _, element := range splitFileArr {
		toDir = toDir + element + "/"
	}
	unzippedContent, err := lib.Unzip(string(fileContent))
	if err != nil {
		log.Fatal(err)
	}
	err = lib.CheckIntegrity(unzippedContent, hash)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(toDir, 666)
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
