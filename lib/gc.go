package lib

import (
	"os"

	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
	"github.com/MSathieu/SimpleVCS/vcstag"
	"github.com/MSathieu/SimpleVCS/vcstree"
)

//GCCommits garbage collects all commits
func GCCommits() error {
	branches, err := vcsbranch.Read()
	if err != nil {
		return err
	}
	tags, err := vcstag.Read()
	if err != nil {
		return err
	}
	for true {
		var removed bool
		commitHashes, err := util.GetAllObjects("commits")
		if err != nil {
			return err
		}
		var referencedCommits []string
		for _, commitHash := range commitHashes {
			commit, err := vcscommit.Get(commitHash)
			if err != nil {
				return err
			}
			if commit.Parent != "" {
				referencedCommits = append(referencedCommits, commit.Parent)
			}
		}
		for _, branch := range branches {
			referencedCommits = append(referencedCommits, branch.Commit.Hash)
		}
		for _, tag := range tags {
			referencedCommits = append(referencedCommits, tag.Commit.Hash)
		}
		for _, commitHash := range commitHashes {
			var exists bool
			for _, referencedCommit := range referencedCommits {
				if commitHash == referencedCommit {
					exists = true
				}
			}
			if !exists {
				removed = true
				err = os.Remove(".svcs/commits/" + commitHash)
				if err != nil {
					return err
				}
			}
		}
		if !removed {
			break
		}
	}
	return nil
}

//GCTrees garbage collects all trees
func GCTrees() error {
	commitHashes, err := util.GetAllObjects("commits")
	if err != nil {
		return err
	}
	treeHashes, err := util.GetAllObjects("trees")
	if err != nil {
		return err
	}
	for _, treeHash := range treeHashes {
		var exists bool
		for _, commitHash := range commitHashes {
			commit, err := vcscommit.Get(commitHash)
			if err != nil {
				return err
			}
			if commit.Tree.Hash == treeHash {
				exists = true
			}
		}
		if !exists {
			err = os.Remove(".svcs/trees/" + treeHash)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//GCFiles garbage collects all files
func GCFiles() error {
	treeHashes, err := util.GetAllObjects("trees")
	if err != nil {
		return err
	}
	fileHashes, err := util.GetAllObjects("files")
	if err != nil {
		return err
	}
	var treeFiles []string
	for _, treeHash := range treeHashes {
		tree, err := vcstree.Get(treeHash)
		if err != nil {
			return err
		}
		for _, treeFile := range tree.Files {
			var exists bool
			for _, treeFileVar := range treeFiles {
				if treeFileVar == treeFile.File.Hash {
					exists = true
				}
			}
			if !exists {
				treeFiles = append(treeFiles, treeFile.File.Hash)
			}
		}
	}
	for _, fileHash := range fileHashes {
		var exists bool
		for _, treeFile := range treeFiles {
			if fileHash == treeFile {
				exists = true
			}
		}
		if !exists {
			err = os.Remove(".svcs/files/" + fileHash)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
