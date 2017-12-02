package cmd

import (
	"errors"
	"fmt"

	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
)

//CreateBranch creates a branch.
func CreateBranch(branch string) error {
	head, err := util.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("can't create branch in detached state")
	}
	headBranch, err := vcsbranch.Get(head.Branch)
	if err != nil {
		return err
	}
	err = vcsbranch.Create(branch, headBranch.Commit.Hash)
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
