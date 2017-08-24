package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Merge(fromBranch string, toBranch string) error {
	if !lib.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	branchesContentByte, _ := ioutil.ReadFile(".svcs/branches.txt")
	branchesContent := string(branchesContentByte)
	fastForward := checkForFastForward(fromBranch, toBranch, branchesContent)
	if fastForward {
		performFastForward(fromBranch, toBranch, branchesContent)
	}
	return nil
}
func checkForFastForward(fromBranch string, toBranch string, branchesContent string) bool {
	var fromSha string
	var toSha string
	for _, line := range strings.Split(branchesContent, "\n") {
		if line == "" {
			continue
		}
		lineArr := strings.Split(line, " ")
		if lineArr[0] == fromBranch {
			fromSha = lineArr[1]
		}
		if lineArr[0] == toBranch {
			toSha = lineArr[1]
		}
	}
	if toSha == "" || fromSha == "" {
		return false
	}
	for currentSha := fromSha; true; {
		if currentSha == toSha {
			return true
		}
		currentSha = lib.GetParent(currentSha)
		if currentSha == "" {
			break
		}
	}
	return false
}
func performFastForward(fromBranch string, toBranch string, branchesContent string) {
	var fromSha string
	branchesArr := strings.Split(branchesContent, "\n")
	var branchesFileContent []string
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == fromBranch {
			fromSha = lineSplit[1]
		}
		if lineSplit[0] != toBranch {
			branchesFileContent = append(branchesFileContent, line)
		}
	}
	branchesFileContent = append(branchesFileContent, toBranch+" "+fromSha)
	branchesFile, _ := os.Create(".svcs/branches.txt")
	for _, line := range branchesFileContent {
		branchesFile.WriteString(line + "\n")
	}
}
