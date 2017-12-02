package pull

import (
	"errors"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
	"github.com/MSathieu/SimpleVCS/vcsbranch"
	"github.com/MSathieu/SimpleVCS/vcscommit"
	"github.com/MSathieu/SimpleVCS/vcstag"
)

//Pull pulls the latest changes.
func Pull(url string, username string, password string) error {
	err := util.ExecHook("prepull")
	if err != nil {
		return err
	}
	if url == "" {
		url, err = util.GetConfig("remote")
		if err != nil {
			return err
		}
	}
	if username == "" {
		username, err = util.GetConfig("username")
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
	files, err := gotils.GetHTTP(url+"/files", authArr)
	if err != nil {
		return err
	}
	filesSplit := strings.Split(files, "\n")
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
	trees, err := gotils.GetHTTP(url+"/trees", authArr)
	if err != nil {
		return err
	}
	treesSplit := strings.Split(trees, "\n")
	for _, treeObj := range treesSplit {
		if treeObj == "" {
			continue
		}
		treeSplit := strings.Split(treeObj, " ")
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
		_, err = vcscommit.SetFiles(filesList)
		if err != nil {
			return err
		}
	}
	commits, err := gotils.GetHTTP(url+"/commits", authArr)
	if err != nil {
		return err
	}
	commitsSplit := strings.Split(commits, "\n")
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
		err = vcsbranch.Update(branchSplit[0], branchSplit[1])
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
		tag, _ := vcstag.Get(tagSplit[0])
		tag.Remove()
		err = vcstag.Create(tagSplit[0], tagSplit[1])
		if err != nil {
			return err
		}
	}
	return util.ExecHook("postpull")
}
