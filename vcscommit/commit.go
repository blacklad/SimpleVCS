package vcscommit

import "github.com/MSathieu/SimpleVCS/vcstree"

//Commit is the commit object.
type Commit struct {
	Author  string
	Time    string
	Parent  string
	Tree    vcstree.Tree
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
