package merge

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//Merge merges two branches.
func Merge(fromBranchString string) error {
	err := util.ExecHook("premerge")
	if err != nil {
		return err
	}
	head, err := util.GetHead()
	if err != nil {
		return err
	}
	if head.Detached {
		return errors.New("cannot merge in detached state")
	}
	fromBranch, err := types.GetBranch(fromBranchString)
	if err != nil {
		return err
	}
	toBranch, err := types.GetBranch(head.Branch)
	if err != nil {
		return err
	}
	fastForward, err := checkForFastForward(fromBranch, toBranch)
	if err != nil {
		return err
	}
	if fastForward {
		err = performFastForward(fromBranch, toBranch)
		if err != nil {
			return err
		}
		err = util.ExecHook("postmerge")
		return err
	}
	parent, err := checkForRecursiveAndGetAncestorSha(fromBranch, toBranch)
	if err != nil {
		return err
	}
	if parent.Hash != "" {
		err = performRecursive(fromBranch, toBranch, parent)
		if err != nil {
			return err
		}
		err = util.ExecHook("postmerge")
		return err
	}
	return errors.New("could not merge")
}
