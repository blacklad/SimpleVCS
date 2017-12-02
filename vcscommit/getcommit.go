package vcscommit

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//Get gets the commit specified by the hash.
func Get(hash string) (types.Commit, error) {
	if hash == "" {
		return types.Commit{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/commits", hash))
	if err != nil {
		return types.Commit{}, err
	}
	fileContent, err := util.Unzip(string(zippedFile))
	if err != nil {
		return types.Commit{}, err
	}
	err = gotils.CheckIntegrity(fileContent, hash)
	if err != nil {
		return types.Commit{}, err
	}
	split := strings.Split(fileContent, "\n")
	var author, time, parent, treeHash, message string
	for _, line := range split {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		switch lineSplit[0] {
		case "author":
			author = lineSplit[1]
		case "time":
			time = lineSplit[1]
		case "parent":
			parent = lineSplit[1]
		case "tree":
			treeHash = lineSplit[1]
		case "message":
			message = lineSplit[1]
		}
	}
	message, err = gotils.Decode(message)
	if err != nil {
		return types.Commit{}, err
	}
	treeObj, err := types.GetTree(treeHash)
	if err != nil {
		return types.Commit{}, err
	}
	return types.Commit{Author: author, Time: time, Parent: parent, Tree: treeObj, Message: message, Hash: hash}, nil
}
