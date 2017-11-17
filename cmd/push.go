package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/lib"
)

var body string

//Push pushes the changes to the server.
func Push(url string, username string, password string) error {
	system, err := gotils.GetHTTP(url+"/system", nil)
	if err != nil {
		return err
	}
	systemSplit := strings.Split(system, " ")
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
		body = body + branch.Name + " " + branch.Commit.Hash + "\n"
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
		body = body + tag.Name + " " + tag.Commit.Hash + "\n"
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
	body = body + info.Name() + " " + gotils.Encode(unzipped) + "\n"
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
	var names []string
	var hashes []string
	for _, file := range tree.Files {
		names = append(names, file.Name)
		hashes = append(hashes, file.File.Hash)
	}
	encodedNames := gotils.Encode(strings.Join(names, " "))
	encodedFiles := gotils.Encode(strings.Join(hashes, " "))
	body = body + info.Name() + " " + encodedNames + " " + encodedFiles + "\n"
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
	body = body +
		commit.Hash + " " +
		commit.Author + " " +
		commit.Parent + " " +
		commit.Tree.Hash + " " +
		commit.Time + " " +
		gotils.Encode(commit.Message) + "\n"
	return nil
}
