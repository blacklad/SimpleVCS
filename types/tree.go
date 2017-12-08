package types

import (
	"strings"

	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/util"
)

//Tree is the tree object.
type Tree struct {
	Hash  string
	Files []TreeFile
}

//TreeFile is the object that has a file and a name.
type TreeFile struct {
	File File
	Name string
}

//GetTree gets a tree
func GetTree(hash string) (Tree, error) {
	if hash == "" {
		return Tree{}, nil
	}
	tree := &util.Tree{}
	util.DB.Where(&Tree{Hash: hash}).First(tree)
	var files []TreeFile
	split := strings.Split(tree.Files, "\n")
	for _, part := range split {
		part, err := gotils.Decode(part)
		if err != nil {
			return Tree{}, err
		}
		partSplit := strings.Split(part, " ")
		filesFile, err := GetFile(partSplit[1])
		if err != nil {
			return Tree{}, err
		}
		files = append(files, TreeFile{File: filesFile, Name: partSplit[0]})
	}
	return Tree{Hash: tree.Hash, Files: files}, nil
}
