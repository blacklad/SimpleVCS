package cmd

import "github.com/MSathieu/SimpleVCS/lib"

//InitRepo inits the repo.
func InitRepo(repoName string, zipped bool, bare bool) error {
	return lib.Init(repoName, zipped, bare)
}
