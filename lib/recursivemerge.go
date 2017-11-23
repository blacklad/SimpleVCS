package lib

import (
	"errors"
)

//CheckForRecursiveAndGetAncestorSha checks if recursive merge is possible and return the ancestor.
func CheckForRecursiveAndGetAncestorSha(fromBranch Branch, toBranch Branch) (Commit, error) {
	var fromCommits []string
	if toBranch.Commit.Hash == "" || fromBranch.Commit.Hash == "" {
		return Commit{}, nil
	}
	for currentCommit := fromBranch.Commit; true; {
		fromCommits = append(fromCommits, currentCommit.Hash)
		currentCommit, err := GetCommit(currentCommit.Parent)
		if err != nil {
			return Commit{}, err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	for currentCommit := toBranch.Commit; true; {
		for _, fromCommit := range fromCommits {
			if fromCommit == currentCommit.Hash {
				return currentCommit, nil
			}
		}
		currentCommit, err := GetCommit(currentCommit.Parent)
		if err != nil {
			return Commit{}, err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	return Commit{}, nil
}

//PerformRecursive performs the recursive merge, run CheckForRecursiveAndGetAncestorSha before running this.
func PerformRecursive(fromBranch Branch, toBranch Branch, parent Commit) error {
	filesArr := parent.GetFiles()
	toChanges := GenerateChange(parent.Tree.Files, toBranch.Commit.Tree.Files)
	fromChanges := GenerateChange(parent.Tree.Files, fromBranch.Commit.Tree.Files)
	for _, toChange := range toChanges {
		for _, fromChange := range fromChanges {
			if toChange.Name == fromChange.Name {
				return errors.New("merge coflict")
			}
		}
	}
	filesArr = ApplyChange(filesArr, toChanges)
	filesArr = ApplyChange(filesArr, fromChanges)
	commitHash, err := CreateCommit("Merged branch "+fromBranch.Name+"into "+toBranch.Name+".", filesArr)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch.Name, commitHash)
	return err
}
