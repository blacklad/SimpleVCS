package cmd

import (
	"errors"
	"fmt"
	"strings"

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
	fmt.Print(strings.Join(lib.ReadBranches(), "\n"))
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
