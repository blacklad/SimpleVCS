package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

//Branch is the branch object
type Branch struct {
	Name   string
	Commit Commit
}

const branchesFile = ".svcs/branches.txt"

//CreateBranch creates the specified branch
func CreateBranch(branch string, sha string) error {
	var branches []Branch
	branchesArr, err := ReadBranches()
	if err != nil {
		return err
	}
	for _, loopBranch := range branchesArr {
		branches = append(branches, loopBranch)
		if loopBranch.Name == branch {
			return errors.New("branch exists")
		}
	}
	commit, err := GetCommit(sha)
	if err != nil {
		return err
	}
	branches = append(branches, Branch{Name: branch, Commit: commit})
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
	var branches []Branch
	branchesArr, err := ReadBranches()
	for _, loopBranch := range branchesArr {
		if loopBranch.Name == branch {
			continue
		}
		branches = append(branches, loopBranch)
	}
	WriteBranches(branches)
	return err
}

//ReadBranches reads the content of branches.txt into an array.
func ReadBranches() ([]Branch, error) {
	branchesContent, err := ioutil.ReadFile(branchesFile)
	if err != nil {
		return nil, err
	}
	var branches []Branch
	for _, line := range strings.Split(string(branchesContent), "\n") {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		commit, err := GetCommit(split[1])
		if err != nil {
			return nil, err
		}
		branches = append(branches, Branch{Name: split[0], Commit: commit})
	}
	return branches, nil
}

//WriteBranches writes the array to branches.txt.
func WriteBranches(branches []Branch) error {
	branchesFile, err := os.Create(branchesFile)
	if err != nil {
		return err
	}
	for _, branch := range branches {
		_, err = branchesFile.WriteString(branch.Name + " " + branch.Commit.Hash + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
