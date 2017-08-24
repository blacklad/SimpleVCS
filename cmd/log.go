package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Log(branch string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	var commits []string
	branches, _ := ioutil.ReadFile(".svcs/branches.txt")
	var lastSha string
	for _, line := range strings.Split(string(branches), "\n") {
		if line == "" {
			continue
		}
		lineArr := strings.Split(line, " ")
		if lineArr[0] == branch {
			lastSha = lineArr[1]
		}
	}
	for currentSha := lastSha; true; {
		commits = append(commits, currentSha)
		currentSha = lib.GetParent(currentSha)
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
	}
	return nil
}
