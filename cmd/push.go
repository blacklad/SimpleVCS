package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

var body string

//Push pushes the changes to the server.
func Push(url string) error {
	systemResponse, err := http.Get(url + "/system")
	if err != nil {
		return err
	}
	systemBytes, err := ioutil.ReadAll(systemResponse.Body)
	if err != nil {
		return err
	}
	systemSplit := strings.Split(string(systemBytes), " ")
	if systemSplit[0] != "simplevcs" {
		return errors.New("unknown server")
	}
	err = filepath.Walk(".svcs/files", visitFilesPush)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/files", "", strings.NewReader(body))
	if err != nil {
		return err
	}
	body = ""
	err = filepath.Walk(".svcs/trees", visitTreesPush)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/trees", "", strings.NewReader(body))
	if err != nil {
		return err
	}
	body = ""
	err = filepath.Walk(".svcs/commits", visitCommitsPush)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/commits", "", strings.NewReader(body))
	if err != nil {
		return err
	}
	body = ""
	branches, err := lib.ReadBranches()
	if err != nil {
		return err
	}
	for _, branch := range branches {
		body = body + branch.Name + branch.Commit.Hash
	}
	_, err = http.Post(url+"/branches", "", strings.NewReader(body))
	if err != nil {
		return err
	}
	body = ""
	tags, err := lib.ReadTags()
	if err != nil {
		return err
	}
	for _, tag := range tags {
		body = body + tag.Name + tag.Commit.Hash
	}
	_, err = http.Post(url+"/tags", "", strings.NewReader(body))
	if err != nil {
		return err
	}
	return nil
}
func visitFilesPush(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	unzipped, err := lib.Unzip(string(content))
	body = body + info.Name() + " " + lib.Encode(unzipped) + "\n"
	return nil
}
func visitTreesPush(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	tree, err := lib.GetTree(info.Name())
	if err != nil {
		return err
	}
	encodedNames := lib.Encode(strings.Join(tree.Names, " "))
	var hashes []string
	for _, file := range tree.Files {
		hashes = append(hashes, file.Hash)
	}
	encodedFiles := lib.Encode(strings.Join(hashes, " "))
	body = body + info.Name() + " " + encodedNames + " " + encodedFiles
	return nil
}
func visitCommitsPush(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	commit, err := lib.GetCommit(info.Name())
	if err != nil {
		return err
	}
	body = body + commit.Hash + " " + commit.Author + " " + commit.Parent + " " + commit.Tree.Hash + " " + commit.Time + " " + lib.Encode(commit.Message)
	return nil
}
