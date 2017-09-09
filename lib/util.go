package lib

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

//VCSExists checks if the .svcs directory exists.
func VCSExists() bool {
	_, err := os.Stat(".svcs")
	if err != nil {
		return false
	}
	return true
}

//GetTime returns the properly formatted date and time.
func GetTime() string {
	return time.Now().Format("20060102150405")
}

//GetTree gets the tree hash of a commit hash.
func GetTree(commitSha string) (string, error) {
	commitFile, err := ioutil.ReadFile(path.Join(".svcs/commits", commitSha+".txt"))
	if err != nil {
		return "", err
	}
	commitArr := strings.Split(string(commitFile), "\n")
	for _, line := range commitArr {
		if line == "" {
			continue
		}
		mapping := strings.Split(line, " ")
		if mapping[0] == "tree" {
			return mapping[1], nil
		}
	}
	return "", nil
}

//GetParent returns the parent sha of the specified commit.
func GetParent(currentSha string) (string, error) {
	currentInfo, err := ioutil.ReadFile(path.Join(".svcs/commits", currentSha+".txt"))
	if err != nil {
		return "", err
	}
	splitFile := strings.Split(string(currentInfo), "\n")
	for _, line := range splitFile {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "parent" {
			if line == "parent " {
				return "", nil
			}
			return splitLine[1], nil
		}
	}
	return "", nil
}

//Zip zips the argument and returns the zipped content.
func Zip(text []byte) string {
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	comp.Write(text)
	comp.Close()
	return compBytes.String()
}

//Unzip unzips the argument and returns the normal content.
func Unzip(text []byte) string {
	var compBytes bytes.Buffer
	compBytes.Write(text)
	comp, _ := gzip.NewReader(&compBytes)
	var outputBytes bytes.Buffer
	outputBytes.ReadFrom(comp)
	comp.Close()
	return outputBytes.String()
}

//ConvertToCommit converts a branch to a hash
func ConvertToCommit(convertFrom string) (string, bool, error) {
	isBranch := false
	commitHash := convertFrom
	branches, err := ReadBranches()
	if err != nil {
		return "", false, err
	}
	for _, branch := range branches {
		if branch == "" {
			continue
		}
		mapping := strings.Split(branch, " ")
		if convertFrom == mapping[0] {
			isBranch = true
			commitHash = mapping[1]
		}
	}
	return commitHash, isBranch, nil
}

//GetHead returns the head.
func GetHead() (string, error) {
	head, err := ioutil.ReadFile(".svcs/head.txt")
	return string(head), err
}

//GetFiles gets the files of a specified commit
func GetFiles(commitHash string) ([]string, error) {
	treeHash, err := GetTree(commitHash)
	if err != nil {
		return nil, err
	}
	filesEntryPath := path.Join(".svcs/trees", treeHash+".txt")
	filesContent, err := ioutil.ReadFile(filesEntryPath)
	if err != nil {
		return nil, err
	}
	unzippedFiles := Unzip(filesContent)
	files := strings.Split(unzippedFiles, "\n")
	return files, nil
}
