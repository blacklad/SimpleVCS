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
	"github.com/MSathieu/SimpleVCS/lib"
)

var wait sync.WaitGroup

//Checkout checks out the specified commit.
func Checkout(commitHash string, noHead bool) error {
	err := lib.ExecHook("precheckout")
	if err != nil {
		return err
	}
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
	files := commit.GetFiles()
	for _, fileEntry := range files {
		if fileEntry == "" {
			continue
		}
		wait.Add(1)
		mapping := strings.Split(fileEntry, " ")
		go concProcessFile(mapping[1], mapping[0])
	}
	if !noHead {
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
	}
	wait.Wait()
	err = lib.ExecHook("postcheckout")
	if err != nil {
		return err
	}
	return nil
}
func concProcessFile(hash string, name string) {
	copyFrom := path.Join(".svcs/files", hash)
	fileContent, err := ioutil.ReadFile(copyFrom)
	if err != nil {
		log.Fatal(err)
	}
	unzippedContent, err := lib.Unzip(string(fileContent))
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
