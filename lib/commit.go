package lib

import (
	"io/ioutil"
	"path"
	"strings"
)

//CommitStruct is the commit object.
type CommitStruct struct {
	Author  string
	Time    string
	Parent  string
	Tree    string
	Message string
}

//GetCommit gets the commit specified by the hash.
func GetCommit(hash string) (CommitStruct, error) {
	file, err := ioutil.ReadFile(path.Join(".svcs/commits", hash+".txt"))
	if err != nil {
		return CommitStruct{}, err
	}
	split := strings.Split(string(file), "\n")
	var author, time, parent, tree, message string
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
			tree = lineSplit[1]
		case "message":
			message = lineSplit[1]
		}
	}
	message, err = Decode(message)
	if err != nil {
		return CommitStruct{}, err
	}
	return CommitStruct{Author: author, Time: time, Parent: parent, Tree: tree, Message: message}, nil
}
