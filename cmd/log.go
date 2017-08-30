package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Log logs all commits.
func Log(branch string) error {
	var commits []string
	var commitMessages []string
	var lastSha string
	branches, err := lib.ReadBranches()
	if err != nil {
		return err
	}
	for _, line := range branches {
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
		message, err := ioutil.ReadFile(path.Join(".svcs/history", currentSha+"_message.txt"))
		if err != nil {
			return err
		}
		commitMessages = append(commitMessages, string(message))
		currentSha, err = lib.GetParent(currentSha)
		if err != nil {
			return err
		}
		if currentSha == "" {
			break
		}
	}
	last := len(commits) - 1
	for i := 0; i < len(commits)/2; i++ {
		commits[i], commits[last-i] = commits[last-i], commits[i]
	}
	for i := 0; i < len(commitMessages)/2; i++ {
		commitMessages[i], commitMessages[last-i] = commitMessages[last-i], commitMessages[i]
	}
	for i, sha := range commits {
		fmt.Println(sha)
		fmt.Println(commitMessages[i])
	}
	return nil
}
