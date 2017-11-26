package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/util"
)

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
	return nil
}

//RemoveStash removes a stash
func RemoveStash(name string) error {
	return os.Remove(".svcs/stashes/" + name)
}
