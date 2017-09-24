package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Merge merges two branches.
func Merge(fromBranchString string) error {
	toBranchString, err := lib.GetHead()
	if err != nil {
		return err
	}
	fromBranch, err := lib.GetBranch(fromBranchString)
	if err != nil {
		return err
	}
	toBranch, err := lib.GetBranch(toBranchString)
	if err != nil {
		return err
	}
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
		err := lib.PerformRecursive(fromBranchString, toBranchString, parentSha)
		return err
	}
	return errors.New("could not merge")
}
