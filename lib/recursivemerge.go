package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func CheckForRecursive(fromBranch string, toBranch string) (string, error) {
	branchesArr, err := ReadBranches()
	if err != nil {
		return "", err
	}
	var fromCommits []string
	var currentFromSha string
	var currentToSha string
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == fromBranch {
			currentFromSha = lineSplit[1]
		}
		if lineSplit[0] == toBranch {
			currentToSha = lineSplit[1]
		}
	}
	if currentToSha == "" || currentFromSha == "" {
		return "", errors.New("branch not found")
	}
	for currentSha := currentFromSha; true; {
		fromCommits = append(fromCommits, currentSha)
		currentSha, err = GetParent(currentSha)
		if err != nil {
			return "", err
		}
		if currentSha == "" {
			break
		}
	}
	for currentSha := currentToSha; true; {
		for _, fromCommit := range fromCommits {
			if fromCommit == currentSha {
				return currentSha, nil
			}
		}
		currentSha, err = GetParent(currentSha)
		if err != nil {
			return "", err
		}
		if currentSha == "" {
			break
		}
	}
	return "", nil
}
func PerformRecursive(fromBranch string, toBranch string, parentSha string) error {
	branchesArr, err := ReadBranches()
	if err != nil {
		return err
	}
	var currentFromSha string
	var currentToSha string
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == fromBranch {
			currentFromSha = lineSplit[1]
		}
		if lineSplit[0] == toBranch {
			currentToSha = lineSplit[1]
		}
	}
	fromFilesByte, err := ioutil.ReadFile(path.Join(".svcs/history", currentFromSha+"_files.txt"))
	if err != nil {
		return err
	}
	fromFiles := string(fromFilesByte)
	toFilesByte, err := ioutil.ReadFile(path.Join(".svcs/history", currentToSha+"_files.txt"))
	if err != nil {
		return err
	}
	toFiles := string(toFilesByte)
	parentFilesByte, err := ioutil.ReadFile(path.Join(".svcs/history", parentSha+"_files.txt"))
	if err != nil {
		return err
	}
	parentFiles := string(parentFilesByte)
	fromFilesArr := strings.Split(fromFiles, "\n")
	toFilesArr := strings.Split(toFiles, "\n")
	parentFilesArr := strings.Split(parentFiles, "\n")
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
	commitMessage, commitHash, err := CreateCommitInfo(GetTime(), toBranch)
	if err != nil {
		return err
	}
	err = CreateCommit(commitMessage, commitHash, "Merged branch "+fromBranch)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch, commitHash)
	if err != nil {
		return err
	}
	filesPath := path.Join(".svcs/history", commitHash+"_files.txt")
	filesFile, err := os.Create(filesPath)
	if err != nil {
		return err
	}
	for _, file := range filesArr {
		_, err = filesFile.WriteString(file + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
