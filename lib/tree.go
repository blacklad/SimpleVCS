package lib

import (
	"io/ioutil"
	"path"
	"strings"
)

//Tree is the tree object.
type Tree struct {
	Hash  string
	Files []File
	Names []string
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
	file := Unzip(string(zippedFile))
	err = CheckIntegrity(file, hash)
	if err != nil {
		return Tree{}, err
	}
	var files []File
	var names []string
	split := strings.Split(file, "\n")
	for _, line := range split {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		filesFile, err := GetFile(lineSplit[0])
		if err != nil {
			return Tree{}, err
		}
		files = append(files, filesFile)
		names = append(names, lineSplit[1])
	}
	return Tree{Hash: hash, Files: files, Names: names}, nil
}
