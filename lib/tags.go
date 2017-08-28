package lib

import (
	"io/ioutil"
	"os"
	"strings"
)

func CreateTag(tag string, sha string) {
	var tags []string
	tagsArr := ReadTags()
	for _, line := range tagsArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		tags = append(tags, line)
		if lineSplit[0] == tag {
			return
		}
	}
	tags = append(tags, tag+" "+sha)
	WriteTags(tags)
}
func RemoveTag(tag string) {
	var tags []string
	tagsArr := ReadTags()
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
	WriteTags(tags)
}
func ReadTags() []string {
	tagsContent, _ := ioutil.ReadFile(".svcs/tags.txt")
	return strings.Split(string(tagsContent), "\n")
}
func WriteTags(tags []string) {
	tagsFile, _ := os.Create(".svcs/tags.txt")
	for _, line := range tags {
		tagsFile.WriteString(line + "\n")
	}
}
