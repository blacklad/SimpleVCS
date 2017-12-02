package checkout

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"

	"github.com/MSathieu/Gotils"
)

var wait sync.WaitGroup

//Checkout checks out the specified commit.
func Checkout(commitHash string, noHead bool) error {
	err := util.ExecHook("precheckout")
	if err != nil {
		return err
	}
	checkoutBranch := commitHash
	currentCommit, isBranch, err := vcsbranch.ConvertToCommit(commitHash)
	if err != nil {
		return err
	}
	commitHash = currentCommit.Hash
	commitObj, err := types.GetCommit(commitHash)
	if err != nil {
		return err
	}
	files := commitObj.GetFiles()
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
	err = checkoutModules()
	if err != nil {
		return err
	}
	return util.ExecHook("postcheckout")
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
