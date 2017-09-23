package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

const branchesFile = ".svcs/branches.txt"

//CreateBranch creates the specified branch
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

//UpdateBranch deletes and then creates the specified branch.
func UpdateBranch(branch string, sha string) error {
	err := RemoveBranch(branch)
	if err != nil {
		return err
	}
	err = CreateBranch(branch, sha)
	return err
}

//RemoveBranch removes branch.
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

//ReadBranches reads the content of branches.txt into an array.
func ReadBranches() ([]string, error) {
	branchesContent, err := ioutil.ReadFile(branchesFile)
	return strings.Split(string(branchesContent), "\n"), err
}

//WriteBranches writes the array to branches.txt.
func WriteBranches(branches []string) error {
	branchesFile, err := os.Create(branchesFile)
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
