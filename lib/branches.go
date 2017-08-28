package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func CreateBranch(branch string, sha string) error {
	var branches []string
	branchesArr, err := ReadBranches()
	if err != nil {
		return err
	}
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		branches = append(branches, line)
		if lineSplit[0] == branch {
			return errors.New("branch exists")
		}
	}
	branches = append(branches, branch+" "+sha)
	err = WriteBranches(branches)
	return err
}
func UpdateBranch(branch string, sha string) error {
	err := RemoveBranch(branch)
	if err != nil {
		return err
	}
	err = CreateBranch(branch, sha)
	return err
}
func RemoveBranch(branch string) error {
	var branches []string
	branchesArr, err := ReadBranches()
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == branch {
			continue
		}
		branches = append(branches, line)
	}
	WriteBranches(branches)
	return err
}
func ReadBranches() ([]string, error) {
	branchesContent, err := ioutil.ReadFile(".svcs/branches.txt")
	return strings.Split(string(branchesContent), "\n"), err
}
func WriteBranches(branches []string) error {
	branchesFile, err := os.Create(".svcs/branches.txt")
	if err != nil {
		return err
	}
	for _, line := range branches {
		_, err = branchesFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
