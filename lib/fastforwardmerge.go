package lib

import (
	"os"
	"strings"
)

func CheckForFastForward(fromBranch string, toBranch string) bool {
	var fromSha string
	var toSha string
	branchesArr := readBranches()
	for _, line := range branchesArr {
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
		currentSha = GetParent(currentSha)
		if currentSha == "" {
			break
		}
	}
	return false
}
func PerformFastForward(fromBranch string, toBranch string) {
	var fromSha string
	branchesArr := readBranches()
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
