package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Merge merges two branches.
func Merge(fromBranchString string) error {
	err := lib.ExecHook("premerge")
	if err != nil {
		return err
	}
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
		err = lib.PerformFastForward(fromBranch, toBranch)
		if err != nil {
			return err
		}
		err = lib.ExecHook("postmerge")
		return err
	}
	parent, err := lib.CheckForRecursiveAndGetAncestorSha(fromBranch, toBranch)
	if err != nil {
		return err
	}
	if parent.Hash != "" {
		err = lib.PerformRecursive(fromBranch, toBranch, parent)
		if err != nil {
			return err
		}
		err = lib.ExecHook("postmerge")
		return err
	}
	return errors.New("could not merge")
}
