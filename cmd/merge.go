package cmd

import (
	"errors"
	"io/ioutil"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Merge(fromBranch string, toBranch string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	branchesContentByte, _ := ioutil.ReadFile(".svcs/branches.txt")
	branchesContent := string(branchesContentByte)
	fastForward := lib.CheckForFastForward(fromBranch, toBranch, branchesContent)
	if fastForward {
		lib.PerformFastForward(fromBranch, toBranch, branchesContent)
		return nil
	}
	parentSha := lib.CheckForRecursive(fromBranch, toBranch, branchesContent)
	if parentSha != "" {
		err := lib.PerformRecursive(fromBranch, toBranch, branchesContent, parentSha)
		return err
	}
	return errors.New("could not merge")
}
