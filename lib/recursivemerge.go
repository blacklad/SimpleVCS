package lib

import (
	"errors"
	"strings"
)

//CheckForRecursiveAndGetAncestorSha checks if recursive merge is possible and return the ancestor sha.
func CheckForRecursiveAndGetAncestorSha(fromBranch string, toBranch string) (string, error) {
	var fromCommits []string
	currentFromCommit, _, err := ConvertToCommit(fromBranch)
	if err != nil {
		return "", err
	}
	currentToCommit, _, err := ConvertToCommit(toBranch)
	if err != nil {
		return "", err
	}
	if currentToCommit.Hash == "" || currentFromCommit.Hash == "" {
		return "", nil
	}
	for currentCommit := currentFromCommit; true; {
		fromCommits = append(fromCommits, currentCommit.Hash)
		if err != nil {
			return "", err
		}
		currentCommit, err = GetCommit(currentCommit.Parent)
		if err != nil {
			return "", err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	for currentCommit := currentToCommit; true; {
		for _, fromCommit := range fromCommits {
			if fromCommit == currentToCommit.Hash {
				return currentCommit.Hash, nil
			}
		}
		if err != nil {
			return "", err
		}
		currentCommit, err = GetCommit(currentCommit.Parent)
		if err != nil {
			return "", err
		}
		if currentCommit.Hash == "" {
			break
		}
	}
	return "", nil
}

//PerformRecursive performs the recursive merge, run CheckForRecursiveAndGetAncestorSha before running this.
func PerformRecursive(fromBranch string, toBranch string, parentSha string) error {
	fromCommit, _, err := ConvertToCommit(fromBranch)
	if err != nil {
		return err
	}
	toCommit, _, err := ConvertToCommit(toBranch)
	if err != nil {
		return err
	}
	fromFilesArr, err := fromCommit.GetFiles()
	if err != nil {
		return err
	}
	toFilesArr, err := toCommit.GetFiles()
	if err != nil {
		return err
	}
	parentCommit, err := GetCommit(parentSha)
	if err != nil {
		return err
	}
	parentFilesArr, err := parentCommit.GetFiles()
	if err != nil {
		return err
	}
	var toChanges []string
	var fromChanges []string
	for _, line := range toFilesArr {
		if line == "" {
			continue
		}
		mapping := strings.Split(line, " ")
		changedStatus := "created"
		for _, parentLine := range parentFilesArr {
			if line == "" {
				continue
			}
			parentMapping := strings.Split(parentLine, " ")
			if parentMapping[0] == mapping[0] {
				if parentMapping[1] == mapping[1] {
					changedStatus = "same"
				} else {
					changedStatus = "changed"
				}
			}
		}
		if changedStatus != "same" {
			toChanges = append(toChanges, changedStatus+" "+mapping[0])
		}
	}
	for _, line := range parentFilesArr {
		if line == "" {
			continue
		}
		mapping := strings.Split(line, " ")
		changedStatus := "deleted"
		for _, toLine := range toFilesArr {
			if line == "" {
				continue
			}
			toMapping := strings.Split(toLine, " ")
			if toMapping[0] == mapping[0] {
				changedStatus = "same"
			}
		}
		if changedStatus != "same" {
			toChanges = append(toChanges, changedStatus+" "+mapping[0])
		}
	}
	for _, line := range fromFilesArr {
		if line == "" {
			continue
		}
		mapping := strings.Split(line, " ")
		changedStatus := "created"
		for _, parentLine := range parentFilesArr {
			if line == "" {
				continue
			}
			parentMapping := strings.Split(parentLine, " ")
			if parentMapping[0] == mapping[0] {
				if parentMapping[1] == mapping[1] {
					changedStatus = "same"
				} else {
					changedStatus = "changed"
				}
			}
		}
		if changedStatus != "same" {
			fromChanges = append(fromChanges, changedStatus+" "+mapping[0])
		}
	}
	for _, line := range parentFilesArr {
		if line == "" {
			continue
		}
		mapping := strings.Split(line, " ")
		changedStatus := "deleted"
		for _, fromLine := range fromFilesArr {
			if line == "" {
				continue
			}
			fromMapping := strings.Split(fromLine, " ")
			if fromMapping[0] == mapping[0] {
				changedStatus = "same"
			}
		}
		if changedStatus != "same" {
			fromChanges = append(fromChanges, changedStatus+" "+mapping[0])
		}
	}
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
				if file == "" {
					continue
				}
				toMapping := strings.Split(file, " ")
				if mapping[1] == toMapping[0] {
					filesArr = append(filesArr, file)
				}
			}
		}
		if mapping[0] == "changed" {
			var updatedLine string
			for _, file := range toFilesArr {
				if file == "" {
					continue
				}
				toMapping := strings.Split(file, " ")
				if mapping[1] == toMapping[0] {
					updatedLine = file
				}
			}
			for i, file := range filesArr {
				if file == "" {
					continue
				}
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr[i] = updatedLine
				}
			}
		}
		if mapping[0] == "deleted" {
			for i, file := range filesArr {
				if file == "" {
					continue
				}
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
				if file == "" {
					continue
				}
				fromMapping := strings.Split(file, " ")
				if mapping[1] == fromMapping[0] {
					filesArr = append(filesArr, file)
				}
			}
		}
		if mapping[0] == "changed" {
			var updatedLine string
			for _, file := range fromFilesArr {
				if file == "" {
					continue
				}
				fromMapping := strings.Split(file, " ")
				if mapping[1] == fromMapping[0] {
					updatedLine = file
				}
			}
			for i, file := range filesArr {
				if file == "" {
					continue
				}
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr[i] = updatedLine
				}
			}
		}
		if mapping[0] == "deleted" {
			for i, file := range filesArr {
				if file == "" {
					continue
				}
				fileMapping := strings.Split(file, " ")
				if fileMapping[0] == mapping[1] {
					filesArr = append(filesArr[:i], filesArr[i+1:]...)
				}
			}
		}
	}
	commitHash, err := CreateCommit("Merged branch "+fromBranch+"into "+toBranch+".", filesArr)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch, commitHash)
	if err != nil {
		return err
	}
	return nil
}
