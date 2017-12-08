package types

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/util"
)

//Tag is the tag object
type Tag struct {
	Name   string
	Commit Commit
}

//GetTag gets the tag.
func GetTag(name string) (Tag, error) {
	tags, err := ReadTags()
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

//CreateTag creates a tag.
func CreateTag(tag string, sha string) error {
	existsTag := &util.Tag{}
	util.DB.Where(&util.Tag{Name: tag}).First(existsTag)
	if existsTag.Name == tag {
		return errors.New("tag existed already")
	}
	util.DB.Create(&util.Tag{Name: tag, Commit: sha})
	return nil
}

//Remove removes the tag.
func (tag Tag) Remove() error {
	util.DB.Delete(&util.Tag{Name: tag.Name})
	return nil
}

//ReadTags reads the tags.txt file into an array
func ReadTags() ([]Tag, error) {
	var tags []util.Tag
	util.DB.Find(tags)
	var returnedTags []Tag
	for _, tag := range tags {
		commit, err := GetCommit(tag.Name)
		if err != nil {
			return nil, err
		}
		returnedTags = append(returnedTags, Tag{Name: tag.Name, Commit: commit})
	}
	return returnedTags, nil
}
