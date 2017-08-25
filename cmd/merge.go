package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Merge(fromBranch string, toBranch string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	fastForward := lib.CheckForFastForward(fromBranch, toBranch)
	if fastForward {
		lib.PerformFastForward(fromBranch, toBranch)
		return nil
	}
	parentSha := lib.CheckForRecursive(fromBranch, toBranch)
	if parentSha != "" {
		err := lib.PerformRecursive(fromBranch, toBranch, parentSha)
		return err
	}
	return errors.New("could not merge")
}
