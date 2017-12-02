package vcscommit

import (
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcstree"
)

//Create creates the commit.
func Create(message string, files []string) (string, error) {
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

func createInfo(tree vcstree.Tree, message string) (Commit, error) {
	head, err := util.GetHead()
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
		Time: gotils.GetTime(),
		Tree: tree, Message: gotils.Encode(message)}
	branchesSplit, err := gotils.SplitFileIntoArr(".svcs/branches.txt")
	if err != nil {
		return Commit{}, err
	}
	for _, line := range branchesSplit {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		if split[0] == head.Branch {
			commit.Parent = split[1]
		}
	}
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
	err := createFile(info, commit.Hash)
	return commit.Hash, err
}

//SetFiles creates a tree.
func SetFiles(files []string) (vcstree.Tree, error) {
	content := strings.Join(files, "\n")
	hash := gotils.GetChecksum(content)
	treeFile, err := os.Create(path.Join(".svcs/trees", hash))
	if err != nil {
		return vcstree.Tree{}, err
	}
	zippedContent, err := util.Zip(content)
	if err != nil {
		return vcstree.Tree{}, err
	}
	_, err = treeFile.WriteString(zippedContent)
	if err != nil {
		return vcstree.Tree{}, nil
	}
	tree, err := vcstree.Get(hash)
	return tree, err
}
