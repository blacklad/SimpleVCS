package lib

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
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
	commit, err := GetCommit(commitSha)
	if err != nil {
		return "", err
	}
	return commit.Tree, nil
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
	filesEntryPath := path.Join(".svcs/trees", treeHash)
	filesContent, err := ioutil.ReadFile(filesEntryPath)
	if err != nil {
		return nil, err
	}
	unzippedFiles := Unzip(filesContent)
	newHash := sha1.Sum([]byte(unzippedFiles))
	if treeHash != fmt.Sprintf("%x", newHash) {
		return nil, errors.New("data has been tampered with")
	}
	files := strings.Split(unzippedFiles, "\n")
	return files, nil
}

//Encode base64 encodes the string.
func Encode(decoded string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(decoded))
	return encoded
}

//Decode decodes the string.
func Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}
