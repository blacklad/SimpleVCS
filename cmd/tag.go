package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func CreateTag(tag string, sha string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	err := lib.CreateTag(tag, sha)
	return err
}
func ListTags() error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	tags, err := lib.ReadTags()
	fmt.Print(strings.Join(tags, "\n"))
	return err
}
func RemoveTag(tag string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	err := lib.RemoveTag(tag)
	return err
}
