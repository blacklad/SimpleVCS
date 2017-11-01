package cmd

import (
	"fmt"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Diff shows the diff between two commits.
func Diff(fromCommitHash string, toCommitHash string) error {
	fromCommit, err := lib.GetCommit(fromCommitHash)
	if err != nil {
		return err
	}
	toCommit, err := lib.GetCommit(toCommitHash)
	if err != nil {
		return err
	}
	fromFiles := fromCommit.GetFiles()
	toFiles := toCommit.GetFiles()
	changes := lib.GenerateChange(fromFiles, toFiles)
	for _, change := range changes {
		fmt.Println(change + ":")
		split := strings.Split(change, " ")
		switch split[0] {
		case "changed":
			fmt.Println("old:")
			for _, file := range fromCommit.Tree.Files {
				if file.Name == split[1] {
					fmt.Println(file.File.Content)
				}
			}
			fmt.Println("new:")
			for _, file := range toCommit.Tree.Files {
				if file.Name == split[1] {
					fmt.Println(file.File.Content)
				}
			}
		case "created":
			for _, file := range toCommit.Tree.Files {
				if file.Name == split[1] {
					fmt.Println(file.File.Content)
				}
			}
		case "deleted":
			for _, file := range fromCommit.Tree.Files {
				if file.Name == split[1] {
					fmt.Println(file.File.Content)
				}
			}
		}
	}
	return nil
}
