package initrepo

import (
	"errors"
	"os"
	"path"

	"github.com/MSathieu/SimpleVCS/util"
)

var exists = util.VCSExists(".svcs")
var settingsPath = ".svcs/settings.txt"

func InitRepo(repoName string) error {
	if exists {
		return errors.New("already initialized")
	}
	os.Mkdir(".svcs", 0700)
	os.Mkdir(".svcs/files", 0700)
	os.Mkdir(".svcs/history", 0700)
	file, _ := os.Create(settingsPath)
	file.Write([]byte("name " + repoName))
	os.Create(path.Join(".svcs", "branches.txt"))
	return nil
}
