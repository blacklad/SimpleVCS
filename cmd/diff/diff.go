package diff

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/types"
)

//Diff shows the diff between two commits.
func Diff(fromCommitHash string, toCommitHash string) error {
	fromCommit, err := types.GetCommit(fromCommitHash)
	if err != nil {
		return err
	}
	toCommit, err := types.GetCommit(toCommitHash)
	if err != nil {
		return err
	}
	changes := types.GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
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
