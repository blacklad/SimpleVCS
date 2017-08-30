package lib

import (
	"strings"
)

// CheckForFastForward checkis if fastforward merge is possible.
func CheckForFastForward(fromBranch string, toBranch string) (bool, error) {
	var fromSha string
	var toSha string
	branchesArr, err := ReadBranches()
	if err != nil {
		return false, err
	}
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
		return false, nil
	}
	for currentSha := fromSha; true; {
		if currentSha == toSha {
			return true, nil
		}
		currentSha, err = GetParent(currentSha)
		if err != nil {
			return false, err
		}
		if currentSha == "" {
			break
		}
	}
	return false, nil
}

//PerformFastForward performs fastforward merge, before calling this you should call CheckForFastforward.
func PerformFastForward(fromBranch string, toBranch string) error {
	var fromSha string
	branchesArr, err := ReadBranches()
	if err != nil {
		return err
	}
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
	err = UpdateBranch(toBranch, fromSha)
	return err
}
