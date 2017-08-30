package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//CreateTag creates a tag.
func CreateTag(tag string, sha string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	err := lib.CreateTag(tag, sha)
	return err
}

//ListTags lists all tags.
func ListTags() error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	tags, err := lib.ReadTags()
	fmt.Print(strings.Join(tags, "\n"))
	return err
}

//RemoveTag removes a tag.
func RemoveTag(tag string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	err := lib.RemoveTag(tag)
	return err
}
