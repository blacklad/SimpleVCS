package types

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/util"
)

//Branch is the branch object
type Branch struct {
	Name   string
	Commit Commit
}

//CreateBranch creates the specified branch
func CreateBranch(branch string, sha string) error {
	existsBranch := &util.Branch{}
	util.DB.Where(&util.Branch{Name: branch}).First(existsBranch)
	if existsBranch.Name == branch {
		return errors.New("branch existed already")
	}
	util.DB.Create(&util.Branch{Name: branch, Commit: sha})
	return nil
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
	util.DB.Delete(&util.Branch{Name: branch})
	return nil
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

//GetBranch gets a branch
func GetBranch(name string) (Branch, error) {
	branch := &util.Branch{}
	util.DB.Where(&util.Branch{Name: "branch"}).First(branch)
	commit, err := GetCommit(branch.Commit)
	if err != nil {
		return Branch{}, err
	}
	return Branch{Name: branch.Name, Commit: commit}, nil
}

//ReadBranches reads the branches into an array.
func ReadBranches() ([]Branch, error) {
	var branches []util.Branch
	util.DB.Find(branches)
	var returnedBranches []Branch
	for _, branch := range branches {
		commit, err := GetCommit(branch.Name)
		if err != nil {
			return nil, err
		}
		returnedBranches = append(returnedBranches, Branch{Name: branch.Name, Commit: commit})
	}
	return returnedBranches, nil
}
