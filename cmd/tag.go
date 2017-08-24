package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func CreateTag(tag string, sha string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
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
func ListTags() error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	tagsContent, _ := ioutil.ReadFile(".svcs/tags.txt")
	tagsArr := strings.Split(string(tagsContent), "\n")
	var tags []string
	for _, line := range tagsArr {
		if line == "" {
			continue
		}
		tags = append(tags, line)
	}
	for _, line := range tags {
		fmt.Println(line)
	}
	return nil
}
