package server

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/vcstag"
)

func pushFiles(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	filesSplit := strings.Split(string(body), "\n")
	for _, fileString := range filesSplit {
		if fileString == "" {
			continue
		}
		fileSplit := strings.Split(fileString, " ")
		_, err := types.GetFile(fileSplit[0])
		if err == nil {
			continue
		}
		decodedFile, err := gotils.Decode(fileSplit[1])
		if err != nil {
			return err
		}
		fileObj := types.File{Hash: fileSplit[0], Content: decodedFile}
		err = fileObj.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

func pushTrees(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	treesSplit := strings.Split(string(body), "\n")
	for _, treeString := range treesSplit {
		if treeString == "" {
			continue
		}
		treeSplit := strings.Split(treeString, " ")
		_, err := types.GetTree(treeSplit[0])
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
		_, err = types.SetFiles(filesList)
		if err != nil {
			return err
		}
	}
	return nil
}

func pushCommits(request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	commitsSplit := strings.Split(string(body), "\n")
	for _, commitObj := range commitsSplit {
		if commitObj == "" {
			continue
		}
		commitSplit := strings.Split(commitObj, " ")
		_, err := types.GetCommit(commitSplit[0])
		if err == nil {
			continue
		}
		commitTree := types.Tree{Hash: commitSplit[3]}
		commitObj := types.Commit{Hash: commitSplit[0], Author: commitSplit[1], Parent: commitSplit[2], Tree: commitTree, Time: commitSplit[4], Message: commitSplit[5]}
		_, err = commitObj.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

func pushBranches(request *http.Request) error {
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
		err = types.UpdateBranch(branchSplit[0], branchSplit[1])
		if err != nil {
			return err
		}
	}
	return nil
}

func pushTags(request *http.Request) error {
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
		tag, _ := vcstag.Get(tagSplit[0])
		tag.Remove()
		err = vcstag.Create(tagSplit[0], tagSplit[1])
		if err != nil {
			return err
		}
	}
	return nil
}
