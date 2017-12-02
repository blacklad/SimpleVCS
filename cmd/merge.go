package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
)

//Merge merges two branches.
func Merge(fromBranchString string) error {
	err := util.ExecHook("premerge")
	if err != nil {
		return err
	}
	head, err := lib.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("cannot merge in detached state")
	}
	fromBranch, err := vcsbranch.Get(fromBranchString)
	if err != nil {
		return err
	}
	toBranch, err := vcsbranch.Get(head.Branch.Name)
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
		err = util.ExecHook("postmerge")
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
		err = util.ExecHook("postmerge")
		return err
	}
	return errors.New("could not merge")
}
