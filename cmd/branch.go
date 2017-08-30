package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//CreateBranch creates a branch.
func CreateBranch(branch string, sha string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	err := lib.CreateBranch(branch, sha)
	return err
}

//ListBranches prints the branches to the output.
func ListBranches() error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	branches, err := lib.ReadBranches()
	if err != nil {
		return err
	}
	fmt.Print(strings.Join(branches, "\n"))
	return nil
}

//RemoveBranch removes a branch.
func RemoveBranch(branch string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	if branch == "master" {
		return errors.New("can't delete master branch")
	}
	err := lib.RemoveBranch(branch)
	return err
}
