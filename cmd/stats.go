package cmd

import (
	"fmt"
	"strconv"

	"github.com/MSathieu/SimpleVCS/lib"
)

//ShowStats displays the repo statistics
func ShowStats() error {
	var branches, tags, commits, contributors int
	branchesArr, err := lib.ReadBranches()
	if err != nil {
		return err
	}
	for range branchesArr {
		branches++
	}
	tagsArr, err := lib.ReadTags()
	if err != nil {
		return err
	}
	for range tagsArr {
		tags++
	}
	fmt.Println("branches " + strconv.Itoa(branches))
	fmt.Println("tags " + strconv.Itoa(tags))
	fmt.Println("commits " + strconv.Itoa(commits))
	fmt.Println("contributors " + strconv.Itoa(contributors))
	return nil
}
