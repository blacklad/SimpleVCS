package patch

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcschange"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//Apply applies a path.
func Apply(filename string) error {
	patchObj, err := Read(filename)
	if err != nil {
		return err
	}
	fromCommit, err := types.GetCommit(patchObj.FromHash)
	if err != nil {
		return err
	}
	var changes []vcschange.Change
	for _, change := range patchObj.Changes {
		split := strings.Split(change, " ")
		hash := ""
		if split[0] != "deleted" {
			content, err := gotils.Decode(split[2])
			if err != nil {
				return err
			}
			hash = gotils.GetChecksum(content)
			fileObj := types.File{Content: content, Hash: hash}
			err = fileObj.Save()
			if err != nil {
				return err
			}
		}
		changes = append(changes, vcschange.Change{Type: split[0], Name: split[1], Hash: hash})
	}
	files := vcschange.ApplyChange(fromCommit.GetFiles(), changes)
	commitHash, err := vcscommit.Create("Applied patch "+filename, files)
	if err != nil {
		return err
	}
	return types.CreateBranch("patch-"+filename, commitHash)
}
