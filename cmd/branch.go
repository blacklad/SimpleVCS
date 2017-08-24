package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func CreateBranch(branch string, sha string) error {
	branchesContent, _ := ioutil.ReadFile(".svcs/branches.txt")
	branchesArr := strings.Split(string(branchesContent), "\n")
	var branches []string
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		branches = append(branches, line)
		if lineSplit[0] == branch {
			return errors.New("branch already exists")
		}
	}
	branches = append(branches, branch+" "+sha)
	branchesFile, _ := os.Create(".svcs/branches.txt")
	for _, line := range branches {
		branchesFile.WriteString(line + "\n")
	}
	return nil
}
