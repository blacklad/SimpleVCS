package vcscommit

import (
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
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

func createInfo(tree types.Tree, message string) (types.Commit, error) {
	head, err := util.GetHead()
	if err != nil {
		return types.Commit{}, err
	}
	username := os.Getenv("SVCS_USERNAME")
	if username == "" {
		currentUser, err := user.Current()
		if err != nil {
			return types.Commit{}, err
		}
		username = currentUser.Name
	}
	username = strings.Fields(username)[0]
	commit := types.Commit{Author: username,
		Time: gotils.GetTime(),
		Tree: tree, Message: gotils.Encode(message)}
	branchesSplit, err := gotils.SplitFileIntoArr(".svcs/branches.txt")
	if err != nil {
		return types.Commit{}, err
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

//SetFiles creates a tree.
func SetFiles(files []string) (types.Tree, error) {
	content := strings.Join(files, "\n")
	hash := gotils.GetChecksum(content)
	treeFile, err := os.Create(path.Join(".svcs/trees", hash))
	if err != nil {
		return types.Tree{}, err
	}
	zippedContent, err := util.Zip(content)
	if err != nil {
		return types.Tree{}, err
	}
	_, err = treeFile.WriteString(zippedContent)
	if err != nil {
		return types.Tree{}, nil
	}
	tree, err := types.GetTree(hash)
	return tree, err
}
