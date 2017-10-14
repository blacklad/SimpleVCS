package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Pull pulls the latest changes.
func Pull(url string) error {
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
	filesResponse, err := http.Get(url + "/files")
	if err != nil {
		return err
	}
	filesBytes, err := ioutil.ReadAll(filesResponse.Body)
	if err != nil {
		return err
	}
	filesSplit := strings.Split(string(filesBytes), "\n")
	for _, file := range filesSplit {
		if file == "" {
			continue
		}
		fileSplit := strings.Split(file, " ")
		_, err := lib.GetFile(fileSplit[0])
		if err == nil {
			continue
		}
		decodedFile, err := lib.Decode(fileSplit[1])
		if err != nil {
			return err
		}
		fileObj := lib.File{Hash: fileSplit[0], Content: decodedFile}
		err = fileObj.Save()
		if err != nil {
			return err
		}
	}
	treesResponse, err := http.Get(url + "/trees")
	if err != nil {
		return err
	}
	treesBytes, err := ioutil.ReadAll(treesResponse.Body)
	if err != nil {
		return err
	}
	treesSplit := strings.Split(string(treesBytes), "\n")
	for _, tree := range treesSplit {
		if tree == "" {
			continue
		}
		treeSplit := strings.Split(tree, " ")
		_, err := lib.GetTree(treeSplit[0])
		if err == nil {
			continue
		}
		decodedNames, err := lib.Decode(treeSplit[1])
		if err != nil {
			return err
		}
		decodedFiles, err := lib.Decode(treeSplit[2])
		if err != nil {
			return err
		}
		filesSplit := strings.Split(decodedFiles, " ")
		namesSplit := strings.Split(decodedNames, " ")
		var filesList []string
		for i := range filesSplit {
			filesList = append(filesList, namesSplit[i]+" "+filesSplit[i])
		}
		_, err = lib.SetFiles(filesList)
		if err != nil {
			return err
		}
	}
	commitsResponse, err := http.Get(url + "/commits")
	if err != nil {
		return err
	}
	commitsBytes, err := ioutil.ReadAll(commitsResponse.Body)
	if err != nil {
		return err
	}
	commitsSplit := strings.Split(string(commitsBytes), "\n")
	for _, commit := range commitsSplit {
		if commit == "" {
			continue
		}
		commitSplit := strings.Split(commit, " ")
		_, err := lib.GetCommit(commitSplit[0])
		if err == nil {
			continue
		}
		commitTree := lib.Tree{Hash: commitSplit[3]}
		commitObj := lib.Commit{Hash: commitSplit[0], Author: commitSplit[1], Parent: commitSplit[2], Tree: commitTree, Time: commitSplit[4], Message: commitSplit[5]}
		_, err = commitObj.Save()
		if err != nil {
			return err
		}
	}
	return nil
}
