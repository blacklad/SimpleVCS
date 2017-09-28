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

//Zip zips the argument and returns the zipped content.
func Zip(text string) string {
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	comp.Write([]byte(text))
	comp.Close()
	return compBytes.String()
}

//Unzip unzips the argument and returns the normal content.
func Unzip(text string) string {
	var compBytes bytes.Buffer
	compBytes.Write([]byte(text))
	comp, _ := gzip.NewReader(&compBytes)
	var outputBytes bytes.Buffer
	outputBytes.ReadFrom(comp)
	comp.Close()
	return outputBytes.String()
}

//ConvertToCommit converts a branch to a hash
func ConvertToCommit(convertFrom string) (Commit, bool, error) {
	isBranch := false
	commitHash := convertFrom
	branches, err := ReadBranches()
	if err != nil {
		return Commit{}, false, err
	}
	for _, branch := range branches {
		if convertFrom == branch.Name {
			isBranch = true
			commitHash = branch.Commit.Hash
		}
	}
	commit, err := GetCommit(commitHash)
	return commit, isBranch, err
}

//GetHead returns the head.
func GetHead() (string, error) {
	head, err := ioutil.ReadFile(".svcs/head.txt")
	return string(head), err
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

//CheckIntegrity checks the integrity.
func CheckIntegrity(content string, hash string) error {
	if hash != GetChecksum(content) {
		return errors.New("data has been tampered with")
	}
	return nil
}

//GetChecksum gets the checksum.
func GetChecksum(data string) string {
	checksum := sha1.Sum([]byte(data))
	return fmt.Sprintf("%x", checksum)
}

//GetConfig gets the config.
func GetConfig(key string) (string, error) {
	file, err := ioutil.ReadFile(".svcs/settings.txt")
	if err != nil {
		return "", err
	}
	split := strings.Split(string(file), "\n")
	for _, line := range split {
		mapping := strings.Split(line, " ")
		if mapping[0] == key {
			return mapping[1], nil
		}
	}
	return "", nil
}
