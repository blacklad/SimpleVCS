package vcsbranch

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
)

//Read reads the content of branches.txt into an array.
func Read() ([]Branch, error) {
	branchesSplit, err := gotils.SplitFileIntoArr(branchesFile)
	if err != nil {
		return nil, err
	}
	var branches []Branch
	for _, line := range branchesSplit {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		var commitObj types.Commit
		if len(split) == 2 {
			commitObj, err = types.GetCommit(split[1])
			if err != nil {
				return nil, err
			}
		} else {
			commitObj = types.Commit{}
		}
		branches = append(branches, Branch{Name: split[0], Commit: commitObj})
	}
	return branches, nil
}
