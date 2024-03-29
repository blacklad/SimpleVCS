package vcslog

import (
	"fmt"
	"time"

	"github.com/MSathieu/SimpleVCS/types"
)

//Log logs all commits.
func Log(branch string) error {
	var commits []types.Commit
	lastCommit, _, err := types.ConvertToCommit(branch)
	if err != nil {
		return err
	}
	for currentCommit := lastCommit; true; {
		commits = append(commits, currentCommit)
		currentCommit, err = types.GetCommit(currentCommit.Parent)
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
	for _, commitObj := range commits {
		time, err := time.Parse("20060102150405", commitObj.Time)
		if err != nil {
			return err
		}
		fmt.Println(commitObj.Hash + " " + commitObj.Author + " " + time.String())
		fmt.Println(commitObj.Message)
	}
	return nil
}
