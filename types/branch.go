package types

import (
	"errors"
	"os"
	"strings"

	"github.com/MSathieu/Gotils"
)

//Branch is the branch object
type Branch struct {
	Name   string
	Commit Commit
}

const branchesFile = ".svcs/branches.txt"

//CreateBranch creates the specified branch
func CreateBranch(branch string, sha string) error {
	branchesArr, err := ReadBranches()
	if err != nil {
		return err
	}
	for _, loopBranch := range branchesArr {
		if loopBranch.Name == branch {
			return errors.New("branch exists")
		}
	}
	commit, err := GetCommit(sha)
	if err != nil {
		return err
	}
	branchesArr = append(branchesArr, Branch{Name: branch, Commit: commit})
	err = WriteBranches(branchesArr)
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

//ConvertToCommit converts a branch to a hash
func ConvertToCommit(convertFrom string) (Commit, bool, error) {
	isBranch := false
	branch, err := GetBranch(convertFrom)
	if err != nil {
		return Commit{}, false, err
	}
	var commit Commit
	if branch.Name != "" {
		isBranch = true
		commit = branch.Commit
	} else {
		commit, err = GetCommit(convertFrom)
		if err != nil {
			return Commit{}, false, err
		}
	}
	return commit, isBranch, nil
}

//ReadBranches reads the content of branches.txt into an array.
func ReadBranches() ([]Branch, error) {
	branchesSplit, err := gotils.SplitFileIntoArr(branchesFile)
	if err != nil {
		return nil, err
	}
	var branches []Branch
	for _, line := range branchesSplit {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		var commitObj Commit
		if len(split) == 2 {
			commitObj, err = GetCommit(split[1])
			if err != nil {
				return nil, err
			}
		} else {
			commitObj = Commit{}
		}
		branches = append(branches, Branch{Name: split[0], Commit: commitObj})
	}
	return branches, nil
}

//GetBranch gets a branch
func GetBranch(name string) (Branch, error) {
	branches, err := ReadBranches()
	if err != nil {
		return Branch{}, err
	}
	for _, branch := range branches {
		if branch.Name == name {
			return branch, nil
		}
	}
	return Branch{}, nil
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
