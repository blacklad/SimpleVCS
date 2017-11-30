package vcstree

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

//Get gets a tree
func Get(hash string) (Tree, error) {
	if hash == "" {
		return Tree{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/trees", hash))
	if err != nil {
		return Tree{}, err
	}
	file, err := util.Unzip(string(zippedFile))
	if err != nil {
		return Tree{}, err
	}
	err = gotils.CheckIntegrity(file, hash)
	if err != nil {
		return Tree{}, err
	}
	var files []File
	split := strings.Split(file, "\n")
	for _, line := range split {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		filesFile, err := vcsfile.GetFile(lineSplit[1])
		if err != nil {
			return Tree{}, err
		}
		files = append(files, File{File: filesFile, Name: lineSplit[0]})
	}
	return Tree{Hash: hash, Files: files}, nil
}
