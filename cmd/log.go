package cmd

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Log logs all commits.
func Log(branch string) error {
	var commits []string
	lastCommit, _, err := lib.ConvertToCommit(branch)
	if err != nil {
		return err
	}
	for currentCommit := lastCommit; true; {
		commits = append(commits, currentCommit.Hash)
		currentCommit, err = lib.GetCommit(currentCommit.Parent)
		if err != nil {
			return err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	last := len(commits) - 1
	for i := 0; i < len(commits)/2; i++ {
		commits[i], commits[last-i] = commits[last-i], commits[i]
	}
	for _, sha := range commits {
		fmt.Println(sha)
		commit, err := lib.GetCommit(sha)
		if err != nil {
			return err
		}
		fmt.Println(commit.Message + "\n")
	}
	return nil
}
