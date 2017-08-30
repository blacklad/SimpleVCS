package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Merge merges two branches.
func Merge(fromBranch string, toBranch string) error {
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
