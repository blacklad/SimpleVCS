package merge

import (
	"github.com/MSathieu/SimpleVCS/types"
)

func checkForFastForward(fromBranch types.Branch, toBranch types.Branch) (bool, error) {
	if toBranch.Commit.Hash == "" || fromBranch.Commit.Hash == "" {
		return false, nil
	}
	for currentCommit := fromBranch.Commit; true; {
		if currentCommit.Hash == toBranch.Commit.Hash {
			return true, nil
		}
		if currentCommit.Parent == "" {
			return false, nil
		}
		parentCommit, err := types.GetCommit(currentCommit.Parent)
		if err != nil {
			return false, err
		}
		currentCommit = parentCommit
	}
	return false, nil
}

func performFastForward(fromBranch types.Branch, toBranch types.Branch) error {
	err := types.UpdateBranch(toBranch.Name, fromBranch.Commit.Hash)
	return err
}
