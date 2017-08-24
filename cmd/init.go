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
	file, _ := os.Create(".svcs/settings.txt")
	file.Write([]byte("name " + repoName))
	os.Create(".svcs/branches.txt")
	return nil
}
