package types

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
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
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/trees", hash))
	if err != nil {
		return Tree{}, err
	}
	fileContent := gotils.UnGZip(string(zippedFile))
	err = gotils.CheckIntegrity(fileContent, hash)
	if err != nil {
		return Tree{}, err
	}
	var files []TreeFile
	split := strings.Split(fileContent, "\n")
	for _, line := range split {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		filesFile, err := GetFile(lineSplit[1])
		if err != nil {
			return Tree{}, err
		}
		files = append(files, TreeFile{File: filesFile, Name: lineSplit[0]})
	}
	return Tree{Hash: hash, Files: files}, nil
}
