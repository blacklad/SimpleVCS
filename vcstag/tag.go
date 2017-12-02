package vcstag

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//Tag is the tag object
type Tag struct {
	Name   string
	Commit types.Commit
}

const tagsFile = ".svcs/tags.txt"

//Get gets the tag.
func Get(name string) (Tag, error) {
	tags, err := Read()
	if err != nil {
		return Tag{}, err
	}
	for _, tag := range tags {
		if tag.Name == name {
			return tag, nil
		}
	}
	return Tag{}, nil
}

//Create creates a tag.
func Create(tag string, sha string) error {
	tagsArr, err := Read()
	if err != nil {
		return err
	}
	for _, loopTag := range tagsArr {
		if loopTag.Name == tag {
			return errors.New("tag already exists")
		}
	}
	commitObj, err := vcscommit.Get(sha)
	if err != nil {
		return err
	}
	tagsArr = append(tagsArr, Tag{Name: tag, Commit: commitObj})
	err = Write(tagsArr)
	return err
}

//Remove removes the tag.
func (tag Tag) Remove() error {
	var tags []Tag
	tagsArr, err := Read()
	if err != nil {
		return err
	}
	for _, loopTag := range tagsArr {
		if loopTag.Name == tag.Name {
			continue
		}
		tags = append(tags, loopTag)
	}
	err = Write(tags)
	return err
}
