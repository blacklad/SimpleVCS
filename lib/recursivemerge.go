package lib

import (
	"errors"
	"strings"
)

//CheckForRecursiveAndGetAncestorSha checks if recursive merge is possible and return the ancestor.
func CheckForRecursiveAndGetAncestorSha(fromBranch Branch, toBranch Branch) (Commit, error) {
	var fromCommits []string
	if toBranch.Commit.Hash == "" || fromBranch.Commit.Hash == "" {
		return Commit{}, nil
	}
	for currentCommit := fromBranch.Commit; true; {
		fromCommits = append(fromCommits, currentCommit.Hash)
		currentCommit, err := GetCommit(currentCommit.Parent)
		if err != nil {
			return Commit{}, err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	for currentCommit := toBranch.Commit; true; {
		for _, fromCommit := range fromCommits {
			if fromCommit == currentCommit.Hash {
				return currentCommit, nil
			}
		}
		currentCommit, err := GetCommit(currentCommit.Parent)
		if err != nil {
			return Commit{}, err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	return Commit{}, nil
}

//PerformRecursive performs the recursive merge, run CheckForRecursiveAndGetAncestorSha before running this.
func PerformRecursive(fromBranch Branch, toBranch Branch, parent Commit) error {
	fromFilesArr := fromBranch.Commit.GetFiles()
	toFilesArr := toBranch.Commit.GetFiles()
	parentFilesArr := parent.GetFiles()
	toChanges := GenerateChange(parentFilesArr, toFilesArr)
	fromChanges := GenerateChange(parentFilesArr, fromFilesArr)
	filesArr := parentFilesArr
	for _, toChange := range toChanges {
		toMapping := strings.Split(toChange, " ")
		for _, fromChange := range fromChanges {
			fromMapping := strings.Split(fromChange, " ")
			if toMapping[1] == fromMapping[1] {
				return errors.New("merge coflict")
			}
		}
	}
	for _, change := range toChanges {
		mapping := strings.Split(change, " ")
		if mapping[0] == "created" {
			for _, file := range toFilesArr {
				toMapping := strings.Split(file, " ")
				if mapping[1] == toMapping[0] {
					filesArr = append(filesArr, file)
				}
			}
		}
		if mapping[0] == "changed" {
			var updatedLine string
			for _, file := range toFilesArr {
				toMapping := strings.Split(file, " ")
				if mapping[1] == toMapping[0] {
					updatedLine = file
				}
			}
			for i, file := range filesArr {
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr[i] = updatedLine
				}
			}
		}
		if mapping[0] == "deleted" {
			for i, file := range filesArr {
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr = append(filesArr[:i], filesArr[i+1:]...)
				}
			}
		}
	}
	for _, change := range fromChanges {
		mapping := strings.Split(change, " ")
		if mapping[0] == "created" {
			for _, file := range fromFilesArr {
				fromMapping := strings.Split(file, " ")
				if mapping[1] == fromMapping[0] {
					filesArr = append(filesArr, file)
				}
			}
		}
		if mapping[0] == "changed" {
			var updatedLine string
			for _, file := range fromFilesArr {
				fromMapping := strings.Split(file, " ")
				if mapping[1] == fromMapping[0] {
					updatedLine = file
				}
			}
			for i, file := range filesArr {
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr[i] = updatedLine
				}
			}
		}
		if mapping[0] == "deleted" {
			for i, file := range filesArr {
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr = append(filesArr[:i], filesArr[i+1:]...)
				}
			}
		}
	}
	commitHash, err := CreateCommit("Merged branch "+fromBranch.Name+"into "+toBranch.Name+".", filesArr)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch.Name, commitHash)
	return err
}
