package cmd

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

//Apply applies a path.
func Apply(filename string) error {
	patch, err := lib.ReadPatch(filename)
	if err != nil {
		return err
	}
	fromCommit, err := vcscommit.Get(patch.FromHash)
	if err != nil {
		return err
	}
	var changes []lib.Change
	for _, change := range patch.Changes {
		split := strings.Split(change, " ")
		hash := ""
		if split[0] != "deleted" {
			content, err := gotils.Decode(split[2])
			if err != nil {
				return err
			}
			hash = gotils.GetChecksum(content)
			fileObj := vcsfile.File{Content: content, Hash: hash}
			err = fileObj.Save()
			if err != nil {
				return err
			}
		}
		changes = append(changes, lib.Change{Type: split[0], Name: split[1], Hash: hash})
	}
	files := lib.ApplyChange(fromCommit.GetFiles(), changes)
	commitHash, err := vcscommit.Create("Applied patch "+filename, files)
	if err != nil {
		return err
	}
	return vcsbranch.Create("patch-"+filename, commitHash)
}
