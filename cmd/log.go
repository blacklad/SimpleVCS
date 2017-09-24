package cmd

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Log logs all commits.
func Log(branch string) error {
	var commits []string
	lastSha, _, err := lib.ConvertToCommit(branch)
	if err != nil {
		return err
	}
	for currentSha := lastSha; true; {
		commits = append(commits, currentSha)
		commit, err := lib.GetCommit(currentSha)
		if err != nil {
			return err
		}
		currentSha = commit.Parent
		if currentSha == "" {
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
