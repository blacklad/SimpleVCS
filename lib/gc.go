package lib

import "os"

//GCCommits garbage collects all commits
func GCCommits() error {
	return nil
}

//GCTrees garbage collects all trees
func GCTrees() error {
	return nil
}

//GCFiles garbage collects all files
func GCFiles() error {
	treeHashes, err := GetAllObjects("trees")
	if err != nil {
		return err
	}
	fileHashes, err := GetAllObjects("files")
	if err != nil {
		return err
	}
	var treeFiles []string
	for _, treeHash := range treeHashes {
		tree, err := GetTree(treeHash)
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
