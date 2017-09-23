package lib

import (
	"crypto/sha1"
	"fmt"
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
	Tree    string
	Message string
}

//GetCommit gets the commit specified by the hash.
func GetCommit(hash string) (Commit, error) {
	file, err := ioutil.ReadFile(path.Join(".svcs/commits", hash+".txt"))
	if err != nil {
		return Commit{}, err
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
		return Commit{}, err
	}
	return Commit{Author: author, Time: time, Parent: parent, Tree: tree, Message: message}, nil
}

//CreateCommit creates the commit.
func CreateCommit(message string, files []string) (string, error) {
	treeHash, err := setFiles(files)
	if err != nil {
		return "", err
	}
	info, err := createCommitInfo(treeHash, message)
	if err != nil {
		return "", err
	}
	sum, err := info.Commit()
	if err != nil {
		return "", err
	}
	return sum, err
}
func createCommitFile(info string, hash string) error {
	infoFile, err := os.Create(path.Join(".svcs/commits", hash+".txt"))
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString(info)
	return err
}

func createCommitInfo(treeHash string, message string) (Commit, error) {
	head, err := GetHead()
	if err != nil {
		return Commit{}, err
	}
	parentSum, _, err := ConvertToCommit(head)
	if err != nil {
		return Commit{}, err
	}
	currentUser, err := user.Current()
	if err != nil {
		return Commit{}, err
	}
	commit := Commit{Author: currentUser.Username, Time: GetTime(), Parent: parentSum, Tree: treeHash, Message: Encode(message)}
	return commit, nil
}

//Commit commits.
func (commit Commit) Commit() (string, error) {
	info := "author " + commit.Author + "\ntime " + commit.Time + "\nparent " + commit.Parent + "\ntree " + commit.Tree + "\nmessage " + commit.Message
	hash := sha1.Sum([]byte(info))
	hashString := fmt.Sprintf("%x", hash)
	err := createCommitFile(info, hashString)
	return hashString, err
}
func setFiles(files []string) (string, error) {
	content := strings.Join(files, "\n")
	hash := sha1.Sum([]byte(content))
	hashString := fmt.Sprintf("%x", hash)
	file, err := os.Create(path.Join(".svcs/trees", hashString))
	if err != nil {
		return "", err
	}
	zippedContent := Zip([]byte(content))
	_, err = file.WriteString(zippedContent)
	return hashString, err
}
