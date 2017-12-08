package initialize

import (
	"errors"
	"os"
)

//Initialize initializes the repo
func Initialize(repoName string) error {
	if repoName == "" {
		return errors.New("you must specify the repo name")
	}
	err := initDirs()
	if err != nil {
		return err
	}
	err = initConfig(repoName)
	if err != nil {
		return err
	}
	branchesFile, err := os.Create(".svcs/branches.txt")
	if err != nil {
		return err
	}
	branchesFile.WriteString("master ")
	_, err = os.Create(".svcs/tags.txt")
	if err != nil {
		return err
	}
	head, err := os.Create(".svcs/head.txt")
	if err != nil {
		return err
	}
	_, err = head.WriteString("master")
	return err
}
