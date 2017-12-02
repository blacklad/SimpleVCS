package vcstag

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
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
		commitObj, err := types.GetCommit(split[1])
		if err != nil {
			return nil, err
		}
		tags = append(tags, Tag{Name: split[0], Commit: commitObj})
	}
	return tags, nil
}
