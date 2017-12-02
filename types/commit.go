package types

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
)

//Commit is the commit object.
type Commit struct {
	Author  string
	Time    string
	Parent  string
	Tree    Tree
	Message string
	Hash    string
}

//GetCommit gets the commit specified by the hash.
func GetCommit(hash string) (Commit, error) {
	if hash == "" {
		return Commit{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/commits", hash))
	if err != nil {
		return Commit{}, err
	}
	fileContent, err := util.Unzip(string(zippedFile))
	if err != nil {
		return Commit{}, err
	}
	err = gotils.CheckIntegrity(fileContent, hash)
	if err != nil {
		return Commit{}, err
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
		return Commit{}, err
	}
	treeObj, err := GetTree(treeHash)
	if err != nil {
		return Commit{}, err
	}
	return Commit{Author: author, Time: time, Parent: parent, Tree: treeObj, Message: message, Hash: hash}, nil
}

//GetFiles gets the files of a specified commit
func (commit Commit) GetFiles() []string {
	var content []string
	for _, fileObj := range commit.Tree.Files {
		content = append(content, fileObj.Name+" "+fileObj.File.Hash)
	}
	return content
}

//Save saves the commit.
func (commit *Commit) Save() (string, error) {
	info := "author " + commit.Author +
		"\ntime " + commit.Time +
		"\nparent " + commit.Parent +
		"\ntree " + commit.Tree.Hash +
		"\nmessage " + commit.Message
	commit.Hash = gotils.GetChecksum(info)
	err := createFile(info, commit.Hash)
	return commit.Hash, err
}
func createFile(info string, hash string) error {
	infoFile, err := os.Create(path.Join(".svcs/commits", hash))
	if err != nil {
		return err
	}
	zipped, err := util.Zip(info)
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString(zipped)
	return err
}
