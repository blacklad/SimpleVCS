package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

//Tag is the tag object
type Tag struct {
	Name   string
	Commit Commit
}

const tagsFile = ".svcs/tags.txt"

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
	tagsArr, err := ReadTags()
	if err != nil {
		return err
	}
	for _, loopTag := range tagsArr {
		if loopTag.Name == tag {
			return errors.New("tag already exists")
		}
	}
	commit, err := GetCommit(sha)
	if err != nil {
		return err
	}
	tagsArr = append(tagsArr, Tag{Name: tag, Commit: commit})
	err = WriteTags(tagsArr)
	return err
}

//Remove removes the tag.
func (tag Tag) Remove() error {
	var tags []Tag
	tagsArr, err := ReadTags()
	if err != nil {
		return err
	}
	for _, loopTag := range tagsArr {
		if loopTag.Name == tag.Name {
			continue
		}
		tags = append(tags, loopTag)
	}
	err = WriteTags(tags)
	return err
}

//ReadTags reads the tags.txt file into an array
func ReadTags() ([]Tag, error) {
	tagsContent, err := ioutil.ReadFile(tagsFile)
	if err != nil {
		return nil, err
	}
	var tags []Tag
	for _, line := range strings.Split(string(tagsContent), "\n") {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		commit, err := GetCommit(split[1])
		if err != nil {
			return nil, err
		}
		tags = append(tags, Tag{Name: split[0], Commit: commit})
	}
	return tags, nil
}

//WriteTags writes an array to tags.txt.
func WriteTags(tags []Tag) error {
	tagsFile, err := os.Create(tagsFile)
	if err != nil {
		return err
	}
	for _, tag := range tags {
		_, err = tagsFile.WriteString(tag.Name + " " + tag.Commit.Hash + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
