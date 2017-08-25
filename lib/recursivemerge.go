package lib

import (
	"io/ioutil"
	"path"
	"strings"
)

func CheckForRecursive(fromBranch string, toBranch string) string {
	branchesArr := readBranches()
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
		return ""
	}
	for currentSha := currentFromSha; true; {
		fromCommits = append(fromCommits, currentSha)
		currentSha = GetParent(currentSha)
		if currentSha == "" {
			break
		}
	}
	for currentSha := currentToSha; true; {
		for _, fromCommit := range fromCommits {
			if fromCommit == currentSha {
				return currentSha
			}
		}
		currentSha = GetParent(currentSha)
		if currentSha == "" {
			break
		}
	}
	return ""
}
func PerformRecursive(fromBranch string, toBranch string, parentSha string) error {
	branchesArr := readBranches()
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
	fromFilesByte, _ := ioutil.ReadFile(path.Join(".svcs/history", currentFromSha+"_files.txt"))
	fromFiles := string(fromFilesByte)
	toFilesByte, _ := ioutil.ReadFile(path.Join(".svcs/history", currentToSha+"_files.txt"))
	toFiles := string(toFilesByte)
	parentFilesByte, _ := ioutil.ReadFile(path.Join(".svcs/history", parentSha+"_files.txt"))
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
	commitMessage, commitHash := CreateMessage(GetTime(), toBranch)
	CreateCommitInfo(commitMessage, commitHash)
	UpdateBranch(toBranch, commitHash)
	return nil
}
