package lib

import (
	"errors"
	"os"
	"strconv"
)

//Init initializes the repo
func Init(repoName string, zipped bool, bare bool) error {
	if repoName == "" {
		return errors.New("you must specify the repo name")
	}
	var dirName string
	if bare {
		dirName = repoName
	} else {
		dirName = ".svcs"
	}
	err := os.Mkdir(dirName, 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/files", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/commits", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/trees", 0700)
	if err != nil {
		return err
	}
	settingsFile, err := os.Create(dirName + "/settings.txt")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("name " + repoName + "\n" + "bare" + "\n")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("zip " + strconv.FormatBool(zipped) + "\n")
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
