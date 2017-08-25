package cmd

import (
	"errors"
	"fmt"

	"github.com/MSathieu/SimpleVCS/lib"
)

func CreateBranch(branch string, sha string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	lib.CreateBranch(branch, sha)
	return nil
}
func ListBranches() error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	fmt.Print(lib.ListBranches())
	return nil
}
func RemoveBranch(branch string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	if branch == "master" {
		return errors.New("cant delete master branch")
	}
	lib.RemoveBranch(branch)
	return nil
}
