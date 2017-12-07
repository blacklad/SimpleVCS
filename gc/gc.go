package gc

import (
	"os"

	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//GC garbage collects all objects
func GC() error {
	err := gcCommits()
	if err != nil {
		return err
	}
	err = gcTrees()
	if err != nil {
		return err
	}
	err = gcFiles()
	return err
}
func gcCommits() error {
	branches, err := types.ReadBranches()
	if err != nil {
		return err
	}
	tags, err := types.ReadTags()
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
			commitObj, err := types.GetCommit(commitHash)
			if err != nil {
				return err
			}
			if commitObj.Parent != "" {
				referencedCommits = append(referencedCommits, commitObj.Parent)
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

func gcTrees() error {
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
			commitObj, err := types.GetCommit(commitHash)
			if err != nil {
				return err
			}
			if commitObj.Tree.Hash == treeHash {
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

func gcFiles() error {
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
		treeObj, err := types.GetTree(treeHash)
		if err != nil {
			return err
		}
		for _, treeFile := range treeObj.Files {
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
