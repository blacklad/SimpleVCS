package checkout

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"

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
	currentCommit, isBranch, err := types.ConvertToCommit(commitHash)
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
		if isBranch {
			util.DB.Where(&util.Config{Name: "head"}).First(&util.Config{}).Update("value", checkoutBranch)
		} else {
			util.DB.Where(&util.Config{Name: "head"}).First(&util.Config{}).Update("value", "DETACHED")
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
	defer wait.Done()
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
	defer newFile.Close()
	newFile.WriteString(file.Content)
	if err != nil {
		log.Fatal(err)
	}
}
