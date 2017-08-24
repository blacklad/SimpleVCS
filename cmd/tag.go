package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func CreateTag(tag string, sha string) error {
	tagsContent, _ := ioutil.ReadFile(".svcs/tags.txt")
	tagsArr := strings.Split(string(tagsContent), "\n")
	var tags []string
	for _, line := range tagsArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		tags = append(tags, line)
		if lineSplit[0] == tag {
			return errors.New("tag already exists")
		}
	}
	tags = append(tags, tag+" "+sha)
	tagsFile, _ := os.Create(".svcs/tags.txt")
	for _, line := range tags {
		tagsFile.WriteString(line + "\n")
	}
	return nil
}
