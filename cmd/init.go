package cmd

import (
	"errors"
	"os"
	"strconv"
)

//InitRepo inits the repo.
func InitRepo(repoName string, zipped bool) error {
	if repoName == "" {
		return errors.New("you must specify the repo name")
	}
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
	err = os.Mkdir(".svcs/hooks", 0700)
	if err != nil {
		return err
	}
	settingsFile, err := os.Create(".svcs/settings.txt")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("name " + repoName + "\n")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("zip " + strconv.FormatBool(zipped) + "\n")
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
	_, err = os.Create(".svcs/hooks/precommit.gake")
	if err != nil {
		return err
	}
	_, err = os.Create(".svcs/hooks/postcommit.gake")
	if err != nil {
		return err
	}
	_, err = os.Create(".svcs/hooks/precheckout.gake")
	if err != nil {
		return err
	}
	_, err = os.Create(".svcs/hooks/postcheckout.gake")
	if err != nil {
		return err
	}
	_, err = os.Create(".svcs/hooks/premerge.gake")
	if err != nil {
		return err
	}_, err = os.Create(".svcs/hooks/postmerge.gake")
	if err != nil {
		return err
	}
	_, err = os.Create(".svcs/ignore.txt")
	return err
}
