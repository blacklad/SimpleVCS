package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/SimpleVCS/util"
)

func Log() error {
	if !util.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	var commits []string
	branches, _ := ioutil.ReadFile(".svcs/branches.txt")
	var lastSha string
	for _, line := range strings.Split(string(branches), "\n") {
		if line == "" {
			continue
		}
		lineArr := strings.Split(line, " ")
		if lineArr[0] == "master" {
			lastSha = lineArr[1]
		}
	}
	for currentSha := lastSha; true; {
		commits = append(commits, currentSha)
		currentSha = getParent(currentSha)
		if currentSha == "" {
			break
		}
	}
	last := len(commits) - 1
	for i := 0; i < len(commits)/2; i++ {
		commits[i], commits[last-i] = commits[last-i], commits[i]
	}
	for _, sha := range commits {
		fmt.Println(sha)
	}
	return nil
}
func getParent(currentSha string) string {
	currentInfo, _ := ioutil.ReadFile(path.Join(".svcs/history", currentSha+".txt"))
	splitFile := strings.Split(string(currentInfo), "\n")
	for _, line := range splitFile {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "parent" {
			if line == "parent " {
				return ""
			}
			return splitLine[1]
		}
	}
	return ""
}
