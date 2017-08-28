package lib

import (
	"io/ioutil"
	"os"
	"strings"
)

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
func ReadTags() ([]string, error) {
	tagsContent, err := ioutil.ReadFile(".svcs/tags.txt")
	return strings.Split(string(tagsContent), "\n"), err
}
func WriteTags(tags []string) error {
	tagsFile, err := os.Create(".svcs/tags.txt")
	if err != nil {
		return err
	}
	for _, line := range tags {
		_, err = tagsFile.WriteString(line + "\n")
		return err
	}
	return nil
}
