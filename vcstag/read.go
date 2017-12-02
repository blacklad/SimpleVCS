package vcstag

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/vcscommit"
)

//Read reads the tags.txt file into an array
func Read() ([]Tag, error) {
	tagsSplit, err := gotils.SplitFileIntoArr(tagsFile)
	if err != nil {
		return nil, err
	}
	var tags []Tag
	for _, line := range tagsSplit {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		commit, err := vcscommit.Get(split[1])
		if err != nil {
			return nil, err
		}
		tags = append(tags, Tag{Name: split[0], Commit: commit})
	}
	return tags, nil
}
