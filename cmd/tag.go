package cmd

import (
	"fmt"

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
	var list string
	for _, tag := range tags {
		list = list + tag.Name + " " + tag.Commit.Hash + "\n"
	}
	fmt.Println(list)
	return err
}

//RemoveTag removes a tag.
func RemoveTag(tag string) error {
	err := lib.RemoveTag(tag)
	return err
}
