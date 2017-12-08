package tag

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/types"
)

//Create creates a tag.
func Create(tag string, sha string) error {
	err := types.CreateTag(tag, sha)
	return err
}

//List lists all tags.
func List() error {
	tags, err := types.ReadTags()
	var list string
	for _, tag := range tags {
		list = list + tag.Name + " " + tag.Commit.Hash + "\n"
	}
	fmt.Println(list)
	return err
}

//Remove removes a tag.
func Remove(name string) error {
	tag, err := types.GetTag(name)
	if err != nil {
		return err
	}
	err = tag.Remove()
	return err
}
