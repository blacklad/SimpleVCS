package cmd

import "github.com/MSathieu/SimpleVCS/initialize"

//InitRepo inits the repo.
func InitRepo(repoName string, zipped bool, bare bool) error {
	return initialize.Initialize(repoName, zipped, bare)
}
