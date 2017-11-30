package lib

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

//PushFiles pushes the files.
func PushFiles(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	filesSplit := strings.Split(string(body), "\n")
	for _, file := range filesSplit {
		if file == "" {
			continue
		}
		fileSplit := strings.Split(file, " ")
		_, err := vcsfile.GetFile(fileSplit[0])
		if err == nil {
			continue
		}
		decodedFile, err := gotils.Decode(fileSplit[1])
		if err != nil {
			return err
		}
		fileObj := vcsfile.File{Hash: fileSplit[0], Content: decodedFile}
		err = fileObj.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

//PushTrees pushes the trees.
func PushTrees(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	treesSplit := strings.Split(string(body), "\n")
	for _, tree := range treesSplit {
		if tree == "" {
			continue
		}
		treeSplit := strings.Split(tree, " ")
		_, err := GetTree(treeSplit[0])
		if err == nil {
			continue
		}
		decodedNames, err := gotils.Decode(treeSplit[1])
		if err != nil {
			return err
		}
		decodedFiles, err := gotils.Decode(treeSplit[2])
		if err != nil {
			return err
		}
		filesSplit := strings.Split(decodedFiles, " ")
		namesSplit := strings.Split(decodedNames, " ")
		var filesList []string
		for i := range filesSplit {
			filesList = append(filesList, namesSplit[i]+" "+filesSplit[i])
		}
		_, err = SetFiles(filesList)
		if err != nil {
			return err
		}
	}
	return nil
}

//PushCommits pushes the commits.
func PushCommits(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	commitsSplit := strings.Split(string(body), "\n")
	for _, commit := range commitsSplit {
		if commit == "" {
			continue
		}
		commitSplit := strings.Split(commit, " ")
		_, err := GetCommit(commitSplit[0])
		if err == nil {
			continue
		}
		commitTree := Tree{Hash: commitSplit[3]}
		commitObj := Commit{Hash: commitSplit[0], Author: commitSplit[1], Parent: commitSplit[2], Tree: commitTree, Time: commitSplit[4], Message: commitSplit[5]}
		_, err = commitObj.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

//PushBranches pushes the branches
func PushBranches(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	branchesSplit := strings.Split(string(body), "\n")
	for _, branch := range branchesSplit {
		if branch == "" {
			continue
		}
		branchSplit := strings.Split(branch, " ")
		err = UpdateBranch(branchSplit[0], branchSplit[1])
		if err != nil {
			return err
		}
	}
	return nil
}

//PushTags pushes the tags
func PushTags(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	tagsSplit := strings.Split(string(body), "\n")
	for _, tag := range tagsSplit {
		if tag == "" {
			continue
		}
		tagSplit := strings.Split(tag, " ")
		tag, _ := GetTag(tagSplit[0])
		tag.Remove()
		err = CreateTag(tagSplit[0], tagSplit[1])
		if err != nil {
			return err
		}
	}
	return nil
}
