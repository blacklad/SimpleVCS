package lib

import (
	"io/ioutil"
	"os"
	"strings"
)

func CreateBranch(branch string, sha string) {
	var branches []string
	branchesArr := ReadBranches()
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		branches = append(branches, line)
		if lineSplit[0] == branch {
			return
		}
	}
	branches = append(branches, branch+" "+sha)
	WriteBranches(branches)
}
func UpdateBranch(branch string, sha string) {
	branchesArr := ReadBranches()
	var branchesFileContent []string
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == branch {
			branchesFileContent = append(branchesFileContent, branch+" "+sha)
			continue
		}
		branchesFileContent = append(branchesFileContent, line)
	}
	WriteBranches(branchesFileContent)
}
func RemoveBranch(branch string) {
	var branches []string
	branchesArr := ReadBranches()
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == branch {
			continue
		}
		branches = append(branches, line)
	}
	WriteBranches(branches)
}
func ReadBranches() []string {
	branchesContent, _ := ioutil.ReadFile(".svcs/branches.txt")
	return strings.Split(string(branchesContent), "\n")
}
func WriteBranches(branches []string) {
	branchesFile, _ := os.Create(".svcs/branches.txt")
	for _, line := range branches {
		branchesFile.WriteString(line + "\n")
	}
}
