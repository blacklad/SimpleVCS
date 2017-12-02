package cmd

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/vcstag"
)

//CreateTag creates a tag.
func CreateTag(tag string, sha string) error {
	err := vcstag.Create(tag, sha)
	return err
}

//ListTags lists all tags.
func ListTags() error {
	tags, err := vcstag.Read()
	var list string
	for _, tag := range tags {
		list = list + tag.Name + " " + tag.Commit.Hash + "\n"
	}
	fmt.Println(list)
	return err
}

//RemoveTag removes a tag.
func RemoveTag(name string) error {
	tag, err := vcstag.Get(name)
	if err != nil {
		return err
	}
	err = tag.Remove()
	return err
}
