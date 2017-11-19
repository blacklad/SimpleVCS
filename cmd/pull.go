package cmd

import (
	"errors"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/lib"
)

//Pull pulls the latest changes.
func Pull(url string, username string, password string) error {
	err := lib.ExecHook("prepull")
	if err != nil {
		return err
	}
	url = "http://" + url + ":333"
	authArr := []string{"USERNAME=" + username, "PASSWORD=" + password}
	system, err := gotils.GetHTTP(url+"/system", nil)
	if err != nil {
		return err
	}
	systemSplit := strings.Split(system, " ")
	if systemSplit[0] != "simplevcs" {
		return errors.New("unknown server")
	}
	files, err := gotils.GetHTTP(url+"/files", authArr)
	if err != nil {
		return err
	}
	filesSplit := strings.Split(files, "\n")
	for _, file := range filesSplit {
		if file == "" {
			continue
		}
		fileSplit := strings.Split(file, " ")
		_, err := lib.GetFile(fileSplit[0])
		if err == nil {
			continue
		}
		decodedFile, err := gotils.Decode(fileSplit[1])
		if err != nil {
			return err
		}
		fileObj := lib.File{Hash: fileSplit[0], Content: decodedFile}
		err = fileObj.Save()
		if err != nil {
			return err
		}
	}
	trees, err := gotils.GetHTTP(url+"/trees", authArr)
	if err != nil {
		return err
	}
	treesSplit := strings.Split(trees, "\n")
	for _, tree := range treesSplit {
		if tree == "" {
			continue
		}
		treeSplit := strings.Split(tree, " ")
		_, err := lib.GetTree(treeSplit[0])
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
		_, err = lib.SetFiles(filesList)
		if err != nil {
			return err
		}
	}
	commits, err := gotils.GetHTTP(url+"/commits", authArr)
	if err != nil {
		return err
	}
	commitsSplit := strings.Split(commits, "\n")
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
	branches, err := gotils.GetHTTP(url+"/branches", authArr)
	if err != nil {
		return err
	}
	branchesSplit := strings.Split(branches, "\n")
	for _, branch := range branchesSplit {
		if branch == "" {
			continue
		}
		branchSplit := strings.Split(branch, " ")
		err = lib.UpdateBranch(branchSplit[0], branchSplit[1])
		if err != nil {
			return err
		}
	}
	tags, err := gotils.GetHTTP(url+"/tags", authArr)
	if err != nil {
		return err
	}
	tagsSplit := strings.Split(tags, "\n")
	for _, tag := range tagsSplit {
		if tag == "" {
			continue
		}
		tagSplit := strings.Split(tag, " ")
		tag, _ := lib.GetTag(tagSplit[0])
		tag.Remove()
		err = lib.CreateTag(tagSplit[0], tagSplit[1])
		if err != nil {
			return err
		}
	}
	return lib.ExecHook("postpull")
}
