package vcsbranch

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//Branch is the branch object
type Branch struct {
	Name   string
	Commit types.Commit
}

const branchesFile = ".svcs/branches.txt"

//Get gets a branch
func Get(name string) (Branch, error) {
	branches, err := Read()
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

//Create creates the specified branch
func Create(branch string, sha string) error {
	branchesArr, err := Read()
	if err != nil {
		return err
	}
	for _, loopBranch := range branchesArr {
		if loopBranch.Name == branch {
			return errors.New("branch exists")
		}
	}
	commit, err := vcscommit.Get(sha)
	if err != nil {
		return err
	}
	branchesArr = append(branchesArr, Branch{Name: branch, Commit: commit})
	err = Write(branchesArr)
	return err
}

//Update deletes and then creates the specified branch.
func Update(branch string, sha string) error {
	err := Remove(branch)
	if err != nil {
		return err
	}
	err = Create(branch, sha)
	return err
}

//Remove removes branch.
func Remove(branch string) error {
	var branches []Branch
	branchesArr, err := Read()
	for _, loopBranch := range branchesArr {
		if loopBranch.Name == branch {
			continue
		}
		branches = append(branches, loopBranch)
	}
	Write(branches)
	return err
}

//ConvertToCommit converts a branch to a hash
func ConvertToCommit(convertFrom string) (types.Commit, bool, error) {
	isBranch := false
	branch, err := Get(convertFrom)
	if err != nil {
		return types.Commit{}, false, err
	}
	var commit types.Commit
	if branch.Name != "" {
		isBranch = true
		commit = branch.Commit
	} else {
		commit, err = vcscommit.Get(convertFrom)
		if err != nil {
			return types.Commit{}, false, err
		}
	}
	return commit, isBranch, nil
}
