package stats

import (
	"fmt"
	"strconv"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//ShowStats displays the repo statistics.
func ShowStats() error {
	branchesArr, err := types.ReadBranches()
	if err != nil {
		return err
	}
	branches := len(branchesArr)
	tagsArr, err := types.ReadTags()
	if err != nil {
		return err
	}
	tags := len(tagsArr)
	var commitsArr []util.Commit
	util.DB.Find(commitsArr)
	commits := len(commitsArr)
	fmt.Println(strconv.Itoa(branches) + " branches")
	fmt.Println(strconv.Itoa(tags) + " tags")
	fmt.Println(strconv.Itoa(commits) + " commits")
	return nil
}
