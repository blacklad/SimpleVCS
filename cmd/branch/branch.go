package branch

import (
	"errors"
	"fmt"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//CreateBranch creates a branch.
func CreateBranch(branch string) error {
	head := util.GetHead()
	if head.Detached {
		return errors.New("can't create branch in detached state")
	}
	headBranch, err := types.GetBranch(head.Branch)
	if err != nil {
		return err
	}
	err = types.CreateBranch(branch, headBranch.Commit.Hash)
	return err
}

//ListBranches prints the branches to the output.
func ListBranches() error {
	branches, err := types.ReadBranches()
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
	err := types.RemoveBranch(branch)
	return err
}
