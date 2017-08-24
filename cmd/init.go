package cmd

import (
	"errors"
	"os"

	"github.com/MSathieu/SimpleVCS/util"
)

func InitRepo(repoName string) error {
	if util.VCSExists(".svcs") {
		return errors.New("not initialized")
	}
	os.Mkdir(".svcs", 0700)
	os.Mkdir(".svcs/files", 0700)
	os.Mkdir(".svcs/history", 0700)
	settingsFile, _ := os.Create(".svcs/settings.txt")
	settingsFile.WriteString("name " + repoName)
	branchesFile, _ := os.Create(".svcs/branches.txt")
	branchesFile.WriteString("master ")
	os.Create(".svcs/tags.txt")
	return nil
}
