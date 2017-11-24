package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//CheckForRecursiveAndGetAncestorSha checks if recursive merge is possible and return the ancestor.
func CheckForRecursiveAndGetAncestorSha(fromBranch Branch, toBranch Branch) (Commit, error) {
	var fromCommits []string
	if toBranch.Commit.Hash == "" || fromBranch.Commit.Hash == "" {
		return Commit{}, nil
	}
	for currentCommit := fromBranch.Commit; true; {
		fromCommits = append(fromCommits, currentCommit.Hash)
		var err error
		currentCommit, err = GetCommit(currentCommit.Parent)
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
		var err error
		currentCommit, err = GetCommit(currentCommit.Parent)
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
	for toI, toChange := range toChanges {
		for fromI, fromChange := range fromChanges {
			if toChange.Name == fromChange.Name {
				fmt.Println("Merge conflict in " + toChange.Name + ":use from branch or to branch?[from|to]")
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				input = strings.Replace(input, "\r\n", "\n", 1)
				switch input {
				case "from\n":
					toChanges[toI] = Change{}
				case "to\n":
					fromChanges[fromI] = Change{}
				default:
					fmt.Print(input)
					return errors.New("aborted due to wrong input")
				}
			}
		}
	}
	var cleanToChanges []Change
	for _, change := range toChanges {
		if change.Type != "" {
			cleanToChanges = append(cleanToChanges, change)
		}
	}
	toChanges = cleanToChanges
	var cleanFromChanges []Change
	for _, change := range fromChanges {
		if change.Type != "" {
			cleanFromChanges = append(cleanFromChanges, change)
		}
	}
	fromChanges = cleanFromChanges
	filesArr = ApplyChange(filesArr, toChanges)
	filesArr = ApplyChange(filesArr, fromChanges)
	commitHash, err := CreateCommit("Merged branch "+fromBranch.Name+"into "+toBranch.Name+".", filesArr)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch.Name, commitHash)
	return err
}
