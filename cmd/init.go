package cmd

import (
	"os"
)

//InitRepo inits the repo.
func InitRepo(repoName string) error {
	err := os.Mkdir(".svcs", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/files", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/commits", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/trees", 0700)
	if err != nil {
		return err
	}
	settingsFile, err := os.Create(".svcs/settings.txt")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("name " + repoName)
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
	if err != nil {
		return err
	}
	_, err = os.Create(".svcs/ignore.txt")
	return err
}
