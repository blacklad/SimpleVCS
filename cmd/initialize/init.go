package initialize

import (
	"errors"
	"os"
)

//Initialize initializes the repo
func Initialize(repoName string, zipped bool, bare bool) error {
	if repoName == "" {
		return errors.New("you must specify the repo name")
	}
	var dirName string
	if bare {
		dirName = repoName
	} else {
		dirName = ".svcs"
	}
	err := initDirs(dirName)
	if err != nil {
		return err
	}
	err = initConfig(repoName, zipped, bare, dirName)
	if err != nil {
		return err
	}
	branchesFile, err := os.Create(dirName + "/branches.txt")
	if err != nil {
		return err
	}
	branchesFile.WriteString("master ")
	_, err = os.Create(dirName + "/tags.txt")
	if err != nil {
		return err
	}
	head, err := os.Create(dirName + "/head.txt")
	if err != nil {
		return err
	}
	_, err = head.WriteString("master")
	return err
}
