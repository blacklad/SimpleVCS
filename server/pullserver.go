package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
	"github.com/MSathieu/SimpleVCS/vcstag"
	"github.com/MSathieu/SimpleVCS/vcstree"
)

func pullFiles(responseWriter http.ResponseWriter) error {
	response = responseWriter
	return filepath.Walk(".svcs/files", visitFiles)
}
func visitFiles(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	unzipped, err := util.Unzip(string(content))
	fmt.Fprintln(response, info.Name()+" "+gotils.Encode(unzipped))
	return nil
}

func pullTrees(responseWriter http.ResponseWriter) error {
	response = responseWriter
	return filepath.Walk(".svcs/trees", visitTrees)
}

func pullCommits(responseWriter http.ResponseWriter) error {
	response = responseWriter
	return filepath.Walk(".svcs/commits", visitCommits)
}
func visitTrees(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	tree, err := vcstree.Get(info.Name())
	if err != nil {
		return err
	}
	var names []string
	var hashes []string
	for _, file := range tree.Files {
		names = append(names, file.Name)
		hashes = append(hashes, file.File.Hash)
	}
	encodedNames := gotils.Encode(strings.Join(names, " "))
	encodedFiles := gotils.Encode(strings.Join(hashes, " "))
	fmt.Fprintln(response, info.Name()+" "+encodedNames+" "+encodedFiles)
	return nil
}
func visitCommits(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	commitObj, err := vcscommit.Get(info.Name())
	if err != nil {
		return err
	}
	fmt.Fprintln(response, commitObj.Hash+" "+commitObj.Author+" "+commitObj.Parent+" "+commitObj.Tree.Hash+" "+commitObj.Time+" "+gotils.Encode(commitObj.Message))
	return nil
}

func pullBranches(responseWriter http.ResponseWriter) error {
	branches, err := vcsbranch.Read()
	if err != nil {
		return err
	}
	for _, branch := range branches {
		fmt.Fprintln(responseWriter, branch.Name+" "+branch.Commit.Hash)
	}
	return nil
}

func pullTags(responseWriter http.ResponseWriter) error {
	tags, err := vcstag.Read()
	if err != nil {
		return err
	}
	for _, tag := range tags {
		fmt.Fprintln(responseWriter, tag.Name+" "+tag.Commit.Hash)
	}
	return nil
}
