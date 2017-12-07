package merge

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcschange"
)

func checkForRecursiveAndGetAncestorSha(fromBranch types.Branch, toBranch types.Branch) (types.Commit, error) {
	var fromCommits []string
	if toBranch.Commit.Hash == "" || fromBranch.Commit.Hash == "" {
		return types.Commit{}, nil
	}
	for currentCommit := fromBranch.Commit; true; {
		fromCommits = append(fromCommits, currentCommit.Hash)
		var err error
		currentCommit, err = types.GetCommit(currentCommit.Parent)
		if err != nil {
			return types.Commit{}, err
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
		currentCommit, err = types.GetCommit(currentCommit.Parent)
		if err != nil {
			return types.Commit{}, err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	return types.Commit{}, nil
}

func performRecursive(fromBranch types.Branch, toBranch types.Branch, parent types.Commit) error {
	filesArr := parent.GetFiles()
	toChanges := vcschange.GenerateChange(parent.Tree.Files, toBranch.Commit.Tree.Files)
	fromChanges := vcschange.GenerateChange(parent.Tree.Files, fromBranch.Commit.Tree.Files)
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
					toChanges[toI] = vcschange.Change{}
				case "to\n":
					fromChanges[fromI] = vcschange.Change{}
				default:
					fmt.Print(input)
					return errors.New("aborted due to wrong input")
				}
			}
		}
	}
	var cleanToChanges []vcschange.Change
	for _, change := range toChanges {
		if change.Type != "" {
			cleanToChanges = append(cleanToChanges, change)
		}
	}
	toChanges = cleanToChanges
	var cleanFromChanges []vcschange.Change
	for _, change := range fromChanges {
		if change.Type != "" {
			cleanFromChanges = append(cleanFromChanges, change)
		}
	}
	fromChanges = cleanFromChanges
	filesArr = vcschange.ApplyChange(filesArr, toChanges)
	filesArr = vcschange.ApplyChange(filesArr, fromChanges)
	commitHash, err := types.CreateCommit("Merged branch "+fromBranch.Name+"into "+toBranch.Name+".", filesArr)
	if err != nil {
		return err
	}
	err = types.UpdateBranch(toBranch.Name, commitHash)
	return err
}
