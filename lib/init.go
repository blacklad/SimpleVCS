package lib

import (
	"errors"
	"os"
	"path"

	"github.com/MSathieu/SimpleVCS/util"
)

func InitRepo(repoName string) error {
	exists := util.VCSExists(".svcs")
	if exists {
		return errors.New("already initialized")
	}
	os.Mkdir(".svcs", 0700)
	file, _ := os.Create(path.Join(".svcs", "settings.txt"))
	file.Write([]byte("name " + repoName))
	os.Create(path.Join(".svcs", "branches.txt"))
	return nil
}
