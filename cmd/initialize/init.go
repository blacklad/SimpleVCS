package initialize

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/util"
)

//Initialize initializes the repo
func Initialize(repoName string) error {
	if repoName == "" {
		return errors.New("you must specify the repo name")
	}
	util.DB.Create(&util.Config{Name: "name", Value: repoName})
	util.DB.Create(&util.Config{Name: "username"})
	util.DB.Create(&util.Config{Name: "remote"})
	util.DB.Create(&util.Branch{Name: "master"})
	util.DB.Create(&util.Config{Name: "head", Value: "master"})
	return nil
}
