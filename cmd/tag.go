package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func CreateTag(tag string, sha string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	lib.CreateTag(tag, sha)
	return nil
}
func ListTags() error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	fmt.Print(strings.Join(lib.ReadTags(), "\n"))
	return nil
}
func RemoveTag(tag string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	lib.RemoveTag(tag)
	return nil
}
