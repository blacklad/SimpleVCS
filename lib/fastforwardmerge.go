package lib

import "github.com/MSathieu/SimpleVCS/vcscommit"

// CheckForFastForward checkis if fastforward merge is possible.
func CheckForFastForward(fromBranch Branch, toBranch Branch) (bool, error) {
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

//PerformFastForward performs fastforward merge, before calling this you should call CheckForFastforward.
func PerformFastForward(fromBranch Branch, toBranch Branch) error {
	err := UpdateBranch(toBranch.Name, fromBranch.Commit.Hash)
	return err
}
