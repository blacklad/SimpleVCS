package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/lib"
	"github.com/MSathieu/SimpleVCS/util"
)

var body string

//Push pushes the changes to the server.
func Push(url string, username string, password string) error {
	err := util.ExecHook("prepush")
	if err != nil {
		return err
	}
	if url == "" {
		url, err = util.GetConfig("remote")
		if err != nil {
			return err
		}
	}
	url = "https://" + url + ":333"
	authArr := []string{"USERNAME=" + username, "PASSWORD=" + password}
	system, err := gotils.GetHTTP(url+"/system", nil)
	if err != nil {
		return err
	}
	systemSplit := strings.Split(system, " ")
	if systemSplit[0] != "simplevcs" {
		return errors.New("unknown server")
	}
	if systemSplit[2] == "multiserver" {
		name, err := util.GetConfig("name")
		if err != nil {
			return err
		}
		url = url + "/" + name
	}
	err = filepath.Walk(".svcs/files", visitFilesPush)
	if err != nil {
		return err
	}
	err = gotils.PostHTTP(url+"/files", body, authArr)
	if err != nil {
		return err
	}
	body = ""
	err = filepath.Walk(".svcs/trees", visitTreesPush)
	if err != nil {
		return err
	}
	err = gotils.PostHTTP(url+"/trees", body, authArr)
	if err != nil {
		return err
	}
	body = ""
	err = filepath.Walk(".svcs/commits", visitCommitsPush)
	if err != nil {
		return err
	}
	err = gotils.PostHTTP(url+"/commits", body, authArr)
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
	err = gotils.PostHTTP(url+"/branches", body, authArr)
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
	err = gotils.PostHTTP(url+"/tags", body, authArr)
	if err != nil {
		return err
	}
	return util.ExecHook("postpush")
}
func visitFilesPush(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	unzipped, err := util.Unzip(string(content))
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
