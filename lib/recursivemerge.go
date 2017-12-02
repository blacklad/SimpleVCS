package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//CheckForRecursiveAndGetAncestorSha checks if recursive merge is possible and return the ancestor.
func CheckForRecursiveAndGetAncestorSha(fromBranch vcsbranch.Branch, toBranch vcsbranch.Branch) (vcscommit.Commit, error) {
	var fromCommits []string
	if toBranch.Commit.Hash == "" || fromBranch.Commit.Hash == "" {
		return vcscommit.Commit{}, nil
	}
	for currentCommit := fromBranch.Commit; true; {
		fromCommits = append(fromCommits, currentCommit.Hash)
		var err error
		currentCommit, err = vcscommit.Get(currentCommit.Parent)
		if err != nil {
			return vcscommit.Commit{}, err
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
		currentCommit, err = vcscommit.Get(currentCommit.Parent)
		if err != nil {
			return vcscommit.Commit{}, err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	return vcscommit.Commit{}, nil
}

//PerformRecursive performs the recursive merge, run CheckForRecursiveAndGetAncestorSha before running this.
func PerformRecursive(fromBranch vcsbranch.Branch, toBranch vcsbranch.Branch, parent vcscommit.Commit) error {
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
	commitHash, err := vcscommit.Create("Merged branch "+fromBranch.Name+"into "+toBranch.Name+".", filesArr)
	if err != nil {
		return err
	}
	err = vcsbranch.Update(toBranch.Name, commitHash)
	return err
}
