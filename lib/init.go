package lib

import (
	"errors"
	"os"
	"path"

	"github.com/MSathieu/SimpleVCS/util"
)

func InitRepo(dir string, repoName string) error {
	exists := util.VCSExists(dir)
	if exists {
		return errors.New("already initialized")
	}
	os.Mkdir(dir, 0700)
	file, _ := os.Create(path.Join(dir, "settings.txt"))
	file.Write([]byte("name " + repoName))
	return nil
}
