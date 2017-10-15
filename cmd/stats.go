package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MSathieu/SimpleVCS/lib"
)

var branches, tags, commits, contributors int

//ShowStats displays the repo statistics.
func ShowStats() error {
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
	err = filepath.Walk(".svcs/commits", visitCommitStats)
	if err != nil {
		return err
	}
	fmt.Println("branches " + strconv.Itoa(branches))
	fmt.Println("tags " + strconv.Itoa(tags))
	fmt.Println("commits " + strconv.Itoa(commits))
	fmt.Println("contributors " + strconv.Itoa(contributors))
	return nil
}
func visitCommitStats(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	_, err = lib.GetCommit(info.Name())
	if err != nil {
		return err
	}
	commits++
	return nil
}
