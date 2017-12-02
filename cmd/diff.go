package cmd

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//Diff shows the diff between two commits.
func Diff(fromCommitHash string, toCommitHash string) error {
	fromCommit, err := vcscommit.Get(fromCommitHash)
	if err != nil {
		return err
	}
	toCommit, err := vcscommit.Get(toCommitHash)
	if err != nil {
		return err
	}
	changes := lib.GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
	for _, change := range changes {
		fmt.Println(change.Type + ":")
		switch change.Type {
		case "changed":
			fmt.Println("old:")
			for _, file := range fromCommit.Tree.Files {
				if file.Name == change.Name {
					fmt.Println(file.File.Content)
				}
			}
			fmt.Println("new:")
			for _, file := range toCommit.Tree.Files {
				if file.Name == change.Name {
					fmt.Println(file.File.Content)
				}
			}
		case "created":
			for _, file := range toCommit.Tree.Files {
				if file.Name == change.Name {
					fmt.Println(file.File.Content)
				}
			}
		case "deleted":
			for _, file := range fromCommit.Tree.Files {
				if file.Name == change.Name {
					fmt.Println(file.File.Content)
				}
			}
		}
	}
	return nil
}
