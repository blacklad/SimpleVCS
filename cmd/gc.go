package cmd

import "github.com/MSathieu/SimpleVCS/lib"

//GarbageCollect cleans the repo up.
func GarbageCollect() error {
	err := lib.GCCommits()
	if err != nil {
		return err
	}
	err = lib.GCTrees()
	if err != nil {
		return err
	}
	return lib.GCFiles()
}
