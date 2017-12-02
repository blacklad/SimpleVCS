package patch

import (
	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcschange"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//Generate generates a patch
func Generate(fromSha string, toSha string, filename string) error {
	fromCommit, err := vcscommit.Get(fromSha)
	if err != nil {
		return err
	}
	toCommit, err := vcscommit.Get(toSha)
	if err != nil {
		return err
	}
	changes := vcschange.GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
	patchObj := Patch{FromHash: fromSha, Changes: []string{}}
	for _, change := range changes {
		changedFile, err := types.GetFile(change.Hash)
		if err != nil {
			return err
		}
		patchObj.Changes = append(patchObj.Changes, change.Type+" "+change.Name+" "+gotils.Encode(changedFile.Content))
	}
	err = patchObj.Save(filename)
	return err
}
