package cmd

import (
	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/vcscommit"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

//GenPatch generates a patch
func GenPatch(fromSha string, toSha string, filename string) error {
	fromCommit, err := vcscommit.Get(fromSha)
	if err != nil {
		return err
	}
	toCommit, err := vcscommit.Get(toSha)
	if err != nil {
		return err
	}
	changes := lib.GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
	patch := lib.Patch{FromHash: fromSha, Changes: []string{}}
	for _, change := range changes {
		changedFile, err := vcsfile.GetFile(change.Hash)
		if err != nil {
			return err
		}
		patch.Changes = append(patch.Changes, change.Type+" "+change.Name+" "+gotils.Encode(changedFile.Content))
	}
	err = patch.Save(filename)
	return err
}
