package lib

import (
	"io/ioutil"
	"os"
	"strings"
)

const tagsFile = ".svcs/tags.txt"

//CreateTag creates a tag.
func CreateTag(tag string, sha string) error {
	var tags []string
	tagsArr, err := ReadTags()
	if err != nil {
		return err
	}
	for _, line := range tagsArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		tags = append(tags, line)
		if lineSplit[0] == tag {
			return nil
		}
	}
	tags = append(tags, tag+" "+sha)
	err = WriteTags(tags)
	return err
}

//RemoveTag removes the tag.
func RemoveTag(tag string) error {
	var tags []string
	tagsArr, err := ReadTags()
	if err != nil {
		return err
	}
	for _, line := range tagsArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == tag {
			continue
		}
		tags = append(tags, line)
	}
	err = WriteTags(tags)
	return err
}

//ReadTags reads the tags.txt file into an array
func ReadTags() ([]string, error) {
	tagsContent, err := ioutil.ReadFile(tagsFile)
	return strings.Split(string(tagsContent), "\n"), err
}

//WriteTags writes an array to tags.txt.
func WriteTags(tags []string) error {
	tagsFile, err := os.Create(tagsFile)
	if err != nil {
		return err
	}
	for _, line := range tags {
		_, err = tagsFile.WriteString(line + "\n")
		return err
	}
	return nil
}
