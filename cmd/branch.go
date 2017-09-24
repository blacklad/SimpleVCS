package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//CreateBranch creates a branch.
func CreateBranch(branch string) error {
	currentBranch, err := lib.GetHead()
	if err != nil {
		return err
	}
	if currentBranch == "DETACHED" {
		return errors.New("can't create branch in detached state")
	}
	sha, _, err := lib.ConvertToCommit(currentBranch)
	if err != nil {
		return err
	}
	err = lib.CreateBranch(branch, sha)
	return err
}

//ListBranches prints the branches to the output.
func ListBranches() error {
	branches, err := lib.ReadBranches()
	if err != nil {
		return err
	}
	fmt.Print(strings.Join(branches, "\n"))
	return nil
}

//RemoveBranch removes a branch.
func RemoveBranch(branch string) error {
	if branch == "master" {
		return errors.New("can't delete master branch")
	}
	err := lib.RemoveBranch(branch)
	return err
}
