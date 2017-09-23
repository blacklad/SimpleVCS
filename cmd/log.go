package cmd

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Log logs all commits.
func Log(branch string) error {
	var commits []string
	var commitMessages []string
	lastSha, _, err := lib.ConvertToCommit(branch)
	if err != nil {
		return err
	}
	for currentSha := lastSha; true; {
		commits = append(commits, currentSha)
		message, err := ioutil.ReadFile(path.Join(".svcs/commits", currentSha+"_message.txt"))
		if err != nil {
			return err
		}
		messageString := lib.Unzip(message)
		commitMessages = append(commitMessages, messageString)
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
