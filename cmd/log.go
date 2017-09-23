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
	lastSha, _, err := lib.ConvertToCommit(branch)
	if err != nil {
		return err
	}
	for currentSha := lastSha; true; {
		commits = append(commits, currentSha)
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
	for _, sha := range commits {
		fmt.Println(sha)
		info, err := ioutil.ReadFile(path.Join(".svcs/commits", sha+".txt"))
		if err != nil {
			return err
		}
		infoSplit := strings.Split(string(info), "\n")
		for _, line := range infoSplit {
			if line == "" {
				continue
			}
			lineSplit := strings.Fields(line)
			if lineSplit[0] == "message" {
				decoded, err := lib.Decode(lineSplit[1])
				if err != nil {
					return err
				}
				fmt.Println(decoded + "\n")
			}
		}
	}
	return nil
}
