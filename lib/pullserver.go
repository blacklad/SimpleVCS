package lib

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
	"github.com/MSathieu/SimpleVCS/vcstree"
)

var response http.ResponseWriter

//PullFiles pulls the files.
func PullFiles(responseWriter http.ResponseWriter) error {
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

//PullTrees pulls the trees.
func PullTrees(responseWriter http.ResponseWriter) error {
	response = responseWriter
	return filepath.Walk(".svcs/trees", visitTrees)
}

//PullCommits pulls the commits.
func PullCommits(responseWriter http.ResponseWriter) error {
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
	commit, err := vcscommit.Get(info.Name())
	if err != nil {
		return err
	}
	fmt.Fprintln(response, commit.Hash+" "+commit.Author+" "+commit.Parent+" "+commit.Tree.Hash+" "+commit.Time+" "+gotils.Encode(commit.Message))
	return nil
}

//PullBranches pulls the branches
func PullBranches(responseWriter http.ResponseWriter) error {
	branches, err := vcsbranch.Read()
	if err != nil {
		return err
	}
	for _, branch := range branches {
		fmt.Fprintln(responseWriter, branch.Name+" "+branch.Commit.Hash)
	}
	return nil
}

//PullTags pulls the tags
func PullTags(responseWriter http.ResponseWriter) error {
	tags, err := ReadTags()
	if err != nil {
		return err
	}
	for _, tag := range tags {
		fmt.Fprintln(responseWriter, tag.Name+" "+tag.Commit.Hash)
	}
	return nil
}
