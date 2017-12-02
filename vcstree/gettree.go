package vcstree

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

//Get gets a tree
func Get(hash string) (types.Tree, error) {
	if hash == "" {
		return types.Tree{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/trees", hash))
	if err != nil {
		return types.Tree{}, err
	}
	fileContent, err := util.Unzip(string(zippedFile))
	if err != nil {
		return types.Tree{}, err
	}
	err = gotils.CheckIntegrity(fileContent, hash)
	if err != nil {
		return types.Tree{}, err
	}
	var files []types.TreeFile
	split := strings.Split(fileContent, "\n")
	for _, line := range split {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		filesFile, err := vcsfile.GetFile(lineSplit[1])
		if err != nil {
			return types.Tree{}, err
		}
		files = append(files, types.TreeFile{File: filesFile, Name: lineSplit[0]})
	}
	return types.Tree{Hash: hash, Files: files}, nil
}
