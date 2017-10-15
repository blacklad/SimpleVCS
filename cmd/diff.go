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
	fromFiles, err := fromCommit.GetFiles()
	if err != nil {
		return err
	}
	toFiles, err := toCommit.GetFiles()
	if err != nil {
		return err
	}
	var changes []string
	for _, line := range fromFiles {
		change := "deleted"
		changed := true
		split := strings.Split(line, " ")
		for _, toLine := range toFiles {
			toSplit := strings.Split(toLine, " ")
			if split[0] == toSplit[0] {
				if split[1] == toSplit[1] {
					changed = false
				} else {
					change = "changed"
				}
			}
		}
		if changed {
			changes = append(changes, change+" "+split[0])
		}
	}
	for _, toLine := range toFiles {
		change := "created"
		changed := true
		toSplit := strings.Split(toLine, " ")
		for _, line := range fromFiles {
			split := strings.Split(line, " ")
			if split[0] == toSplit[0] {
				changed = false
			}
		}
		if changed {
			changes = append(changes, change+" "+toSplit[0])
		}
	}
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
