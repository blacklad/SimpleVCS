package lib

import (
	"strings"
)

func CheckForFastForward(fromBranch string, toBranch string) bool {
	var fromSha string
	var toSha string
	branchesArr := ReadBranches()
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
	branchesArr := ReadBranches()
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == fromBranch {
			fromSha = lineSplit[1]
			break
		}
	}
	UpdateBranch(toBranch, fromSha)
}
