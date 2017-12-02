package types

import (
	"os"
	"path"

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
