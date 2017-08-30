package cmd

import (
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//CreateTag creates a tag.
func CreateTag(tag string, sha string) error {
	err := lib.CreateTag(tag, sha)
	return err
}

//ListTags lists all tags.
func ListTags() error {
	tags, err := lib.ReadTags()
	fmt.Print(strings.Join(tags, "\n"))
	return err
}

//RemoveTag removes a tag.
func RemoveTag(tag string) error {
	err := lib.RemoveTag(tag)
	return err
}
