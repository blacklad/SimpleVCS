package lib

import (
	"crypto/sha1"
	"errors"
	"fmt"
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
	newSha := sha1.Sum([]byte(file))
	if hash != fmt.Sprintf("%x", newSha) {
		return Tree{}, errors.New("data has been tampered with")
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
