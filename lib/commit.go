package lib

import (
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcstree"
)

//Commit is the commit object.
type Commit struct {
	Author  string
	Time    string
	Parent  string
	Tree    vcstree.Tree
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
	file, err := util.Unzip(string(zippedFile))
	if err != nil {
		return Commit{}, err
	}
	err = gotils.CheckIntegrity(file, hash)
	if err != nil {
		return Commit{}, err
	}
	split := strings.Split(file, "\n")
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
	tree, err := vcstree.Get(treeHash)
	if err != nil {
		return Commit{}, err
	}
	return Commit{Author: author, Time: time, Parent: parent, Tree: tree, Message: message, Hash: hash}, nil
}

//CreateCommit creates the commit.
func CreateCommit(message string, files []string) (string, error) {
	tree, err := SetFiles(files)
	if err != nil {
		return "", err
	}
	info, err := createCommitInfo(tree, message)
	if err != nil {
		return "", err
	}
	sum, err := info.Save()
	if err != nil {
		return "", err
	}
	return sum, err
}
func createCommitFile(info string, hash string) error {
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

func createCommitInfo(tree vcstree.Tree, message string) (Commit, error) {
	head, err := GetHead()
	if err != nil {
		return Commit{}, err
	}
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
		Time:   gotils.GetTime(),
		Parent: head.Branch.Commit.Hash,
		Tree:   tree, Message: gotils.Encode(message)}
	return commit, nil
}

//Save saves the commit.
func (commit *Commit) Save() (string, error) {
	info := "author " + commit.Author +
		"\ntime " + commit.Time +
		"\nparent " + commit.Parent +
		"\ntree " + commit.Tree.Hash +
		"\nmessage " + commit.Message
	commit.Hash = gotils.GetChecksum(info)
	err := createCommitFile(info, commit.Hash)
	return commit.Hash, err
}

//SetFiles creates a tree.
func SetFiles(files []string) (vcstree.Tree, error) {
	content := strings.Join(files, "\n")
	hash := gotils.GetChecksum(content)
	file, err := os.Create(path.Join(".svcs/trees", hash))
	if err != nil {
		return vcstree.Tree{}, err
	}
	zippedContent, err := util.Zip(content)
	if err != nil {
		return vcstree.Tree{}, err
	}
	_, err = file.WriteString(zippedContent)
	if err != nil {
		return vcstree.Tree{}, nil
	}
	tree, err := vcstree.Get(hash)
	return tree, err
}

//GetFiles gets the files of a specified commit
func (commit Commit) GetFiles() []string {
	var content []string
	for _, file := range commit.Tree.Files {
		content = append(content, file.Name+" "+file.File.Hash)
	}
	return content
}
