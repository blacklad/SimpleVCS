package lib

import (
	"io/ioutil"
	"os"
	"strings"
)

func CreateBranch(branch string, sha string) {
	var branches []string
	branchesArr := readBranches()
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
	writebranches(branches)
}
func UpdateBranch(branch string, sha string) {
	branchesArr := readBranches()
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
	writebranches(branchesFileContent)
}
func RemoveBranch(branch string) {
	var branches []string
	branchesArr := readBranches()
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
	writebranches(branches)
}
func ListBranches() string {
	branchesArr := readBranches()
	var branches string
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		branches = branches + line + "\n"
	}
	return branches
}
func readBranches() []string {
	branchesContent, _ := ioutil.ReadFile(".svcs/branches.txt")
	return strings.Split(string(branchesContent), "\n")
}
func writebranches(branches []string) {
	branchesFile, _ := os.Create(".svcs/branches.txt")
	for _, line := range branches {
		branchesFile.WriteString(line + "\n")
	}
}
