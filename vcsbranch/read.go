package vcsbranch

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/vcscommit"
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
		var commit vcscommit.Commit
		if len(split) == 2 {
			commit, err = vcscommit.Get(split[1])
			if err != nil {
				return nil, err
			}
		} else {
			commit = vcscommit.Commit{}
		}
		branches = append(branches, Branch{Name: split[0], Commit: commit})
	}
	return branches, nil
}
