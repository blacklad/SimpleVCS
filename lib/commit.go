package lib

import (
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
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
	file, err := Unzip(string(zippedFile))
	if err != nil {
		return Commit{}, err
	}
	err = CheckIntegrity(file, hash)
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
	message, err = Decode(message)
	if err != nil {
		return Commit{}, err
	}
	tree, err := GetTree(treeHash)
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
	zipped, err := Zip(info)
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString(zipped)
	return err
}

func createCommitInfo(tree Tree, message string) (Commit, error) {
	head, err := GetHead()
	if err != nil {
		return Commit{}, err
	}
	parent, _, err := ConvertToCommit(head.Branch.Name)
	if err != nil {
		return Commit{}, err
	}
	currentUser, err := user.Current()
	if err != nil {
		return Commit{}, err
	}
	commit := Commit{Author: currentUser.Username, Time: GetTime(), Parent: parent.Hash, Tree: tree, Message: Encode(message)}
	return commit, nil
}

//Save saves the commit.
func (commit Commit) Save() (string, error) {
	info := "author " + commit.Author + "\ntime " + commit.Time + "\nparent " + commit.Parent + "\ntree " + commit.Tree.Hash + "\nmessage " + commit.Message
	hash := GetChecksum(info)
	err := createCommitFile(info, hash)
	return hash, err
}

//SetFiles creates a tree.
func SetFiles(files []string) (Tree, error) {
	content := strings.Join(files, "\n")
	hash := GetChecksum(content)
	file, err := os.Create(path.Join(".svcs/trees", hash))
	if err != nil {
		return Tree{}, err
	}
	zippedContent, err := Zip(content)
	if err != nil {
		return Tree{}, err
	}
	_, err = file.WriteString(zippedContent)
	if err != nil {
		return Tree{}, nil
	}
	tree, err := GetTree(hash)
	return tree, err
}

//GetFiles gets the files of a specified commit
func (commit Commit) GetFiles() ([]string, error) {
	var content []string
	for i, file := range commit.Tree.Files {
		content = append(content, commit.Tree.Names[i]+" "+file.Hash)
	}
	return content, nil
}
