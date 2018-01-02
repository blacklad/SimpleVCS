package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

func pullFiles(responseWriter http.ResponseWriter) {
	var files []util.File
	util.DB.Find(&files)
	for _, file := range files {
		fmt.Println(responseWriter, file.Hash+" "+gotils.Encode(file.Content))
	}
}

func pullTrees(responseWriter http.ResponseWriter) error {
	var trees []util.Tree
	util.DB.Find(&trees)
	for _, tree := range trees {
		treeObj, err := types.GetTree(tree.Hash)
		if err != nil {
			return err
		}
		var names []string
		var hashes []string
		for _, fileObj := range treeObj.Files {
			names = append(names, fileObj.Name)
			hashes = append(hashes, fileObj.File.Hash)
		}
		encodedNames := gotils.Encode(strings.Join(names, " "))
		encodedFiles := gotils.Encode(strings.Join(hashes, " "))
		fmt.Fprintln(responseWriter, tree.Hash+" "+encodedNames+" "+encodedFiles)
	}
	return nil
}

func pullCommits(responseWriter http.ResponseWriter) error {
	var commits []util.Commit
	util.DB.Find(&commits)
	for _, commit := range commits {
		fmt.Println(responseWriter, commit.Hash+" "+commit.Author+" "+commit.Parent+" "+commit.Tree+" "+commit.Time+" "+gotils.Encode(commit.Message))
	}
	return nil
}

func pullBranches(responseWriter http.ResponseWriter) error {
	branches, err := types.ReadBranches()
	if err != nil {
		return err
	}
	for _, branch := range branches {
		fmt.Fprintln(responseWriter, branch.Name+" "+branch.Commit.Hash)
	}
	return nil
}

func pullTags(responseWriter http.ResponseWriter) error {
	tags, err := types.ReadTags()
	if err != nil {
		return err
	}
	for _, tag := range tags {
		fmt.Fprintln(responseWriter, tag.Name+" "+tag.Commit.Hash)
	}
	return nil
}
