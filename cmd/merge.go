package cmd

import (
	"errors"
	"io/ioutil"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Merge merges two branches.
func Merge(fromBranch string) error {
	toBranchBytes, err := ioutil.ReadFile(".svcs/head.txt")
	if err != nil {
		return err
	}
	toBranch := string(toBranchBytes)
	fastForward, err := lib.CheckForFastForward(fromBranch, toBranch)
	if err != nil {
		return err
	}
	if fastForward {
		err := lib.PerformFastForward(fromBranch, toBranch)
		return err
	}
	parentSha, err := lib.CheckForRecursiveAndGetAncestorSha(fromBranch, toBranch)
	if err != nil {
		return err
	}
	if parentSha != "" {
		err := lib.PerformRecursive(fromBranch, toBranch, parentSha)
		return err
	}
	return errors.New("could not merge")
}
