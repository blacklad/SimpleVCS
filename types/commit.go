package types

import (
	"os"
	"os/user"
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
	commit := &util.Commit{}
	util.DB.Where(&util.Commit{Hash: hash}).First(commit)
	tree, err := GetTree(commit.Tree)
	if err != nil {
		return Commit{}, err
	}
	return Commit{Author: commit.Author, Time: commit.Time, Parent: commit.Parent, Tree: tree, Message: commit.Message, Hash: commit.Hash}, nil
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
	util.DB.Create(&util.Commit{Hash: commit.Hash, Author: commit.Author, Time: commit.Time, Tree: commit.Tree.Hash, Message: commit.Message})
	return commit.Hash, nil
}

//CreateCommit creates the commit.
func CreateCommit(message string, files []string) (string, error) {
	tree, err := SetFiles(files)
	if err != nil {
		return "", err
	}
	info, err := createInfo(tree, message)
	if err != nil {
		return "", err
	}
	sum, err := info.Save()
	if err != nil {
		return "", err
	}
	return sum, err
}

func createInfo(tree Tree, message string) (Commit, error) {
	head := util.GetHead()
	username := os.Getenv("SVCS_USERNAME")
	if username == "" {
		currentUser, err := user.Current()
		if err != nil {
			return Commit{}, err
		}
		username = currentUser.Name
	}
	username = strings.Fields(username)[0]
	commit := Commit{Author: username,
		Time: gotils.GetTime(),
		Tree: tree, Message: gotils.Encode(message)}
	var branch util.Branch
	util.DB.Where(&util.Branch{Name: head.Branch}).First(&branch)
	commit.Parent = branch.Commit
	return commit, nil
}

//SetFiles creates a tree.
func SetFiles(files []string) (Tree, error) {
	content := strings.Join(files, "\n")
	hash := gotils.GetChecksum(content)
	util.DB.Create(&util.Tree{Hash: hash, Files: content})
	tree, err := GetTree(hash)
	return tree, err
}
