package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcstag"
)

var branches, tags, commits int
var contributors []string

//ShowStats displays the repo statistics.
func ShowStats() error {
	branchesArr, err := types.ReadBranches()
	if err != nil {
		return err
	}
	for range branchesArr {
		branches++
	}
	tagsArr, err := vcstag.Read()
	if err != nil {
		return err
	}
	for range tagsArr {
		tags++
	}
	err = filepath.Walk(".svcs/commits", visitCommitStats)
	if err != nil {
		return err
	}
	fmt.Println("branches " + strconv.Itoa(branches))
	fmt.Println("tags " + strconv.Itoa(tags))
	fmt.Println("commits " + strconv.Itoa(commits))
	fmt.Println("contributors " + strconv.Itoa(len(contributors)))
	return nil
}
func visitCommitStats(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	commitObj, err := types.GetCommit(info.Name())
	if err != nil {
		return err
	}
	var isInContributors = false
	for _, cont := range contributors {
		if cont == commitObj.Author {
			isInContributors = true
			break
		}
	}
	if !isInContributors {
		contributors = append(contributors, commitObj.Author)
	}
	commits++
	return nil
}
