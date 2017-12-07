package patch

import (
	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/types"
)

//Generate generates a patch
func Generate(fromSha string, toSha string, filename string) error {
	fromCommit, err := types.GetCommit(fromSha)
	if err != nil {
		return err
	}
	toCommit, err := types.GetCommit(toSha)
	if err != nil {
		return err
	}
	changes := types.GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
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
