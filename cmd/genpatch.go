package cmd

import (
	"os"

	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/lib"
)

//GenPatch generates a patch
func GenPatch(fromSha string, toSha string, filename string) error {
	fromCommit, err := lib.GetCommit(fromSha)
	if err != nil {
		return err
	}
	toCommit, err := lib.GetCommit(toSha)
	if err != nil {
		return err
	}
	changes := lib.GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
	patchFile := "parent " + fromSha + "\n"
	for _, change := range changes {
		changedFile, err := lib.GetFile(change.Hash)
		if err != nil {
			return err
		}
		patchFile = patchFile + change.Type + " " + change.Name + " " + gotils.Encode(changedFile.Content) + "\n"
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(patchFile)
	return err
}
