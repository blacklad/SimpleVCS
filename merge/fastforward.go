package merge

import (
	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

func checkForFastForward(fromBranch vcsbranch.Branch, toBranch vcsbranch.Branch) (bool, error) {
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
		parentCommit, err := vcscommit.Get(currentCommit.Parent)
		if err != nil {
			return false, err
		}
		currentCommit = parentCommit
	}
	return false, nil
}

func performFastForward(fromBranch vcsbranch.Branch, toBranch vcsbranch.Branch) error {
	err := vcsbranch.Update(toBranch.Name, fromBranch.Commit.Hash)
	return err
}