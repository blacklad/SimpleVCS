package cmd

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Log logs all commits.
func Log(branch string) error {
	var commits []lib.Commit
	lastCommit, _, err := lib.ConvertToCommit(branch)
	if err != nil {
		return err
	}
	for currentCommit := lastCommit; true; {
		commits = append(commits, currentCommit)
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
	for _, commit := range commits {
		fmt.Println(commit.Hash + " " + commit.Author + " " + commit.Time)
		fmt.Println(commit.Message)
	}
	return nil
}
