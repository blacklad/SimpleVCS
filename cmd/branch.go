package cmd

import (
	"errors"
	"fmt"

	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
)

//CreateBranch creates a branch.
func CreateBranch(branch string) error {
	head, err := lib.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("can't create branch in detached state")
	}
	err = vcsbranch.Create(branch, head.Branch.Commit.Hash)
	return err
}

//ListBranches prints the branches to the output.
func ListBranches() error {
	branches, err := vcsbranch.Read()
	var list string
	for _, branch := range branches {
		list = list + branch.Name + " " + branch.Commit.Hash + "\n"
	}
	fmt.Println(list)
	return err
}

//RemoveBranch removes a branch.
func RemoveBranch(branch string) error {
	if branch == "master" {
		return errors.New("can't delete master branch")
	}
	err := vcsbranch.Remove(branch)
	return err
}
