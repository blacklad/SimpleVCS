package lib

import (
	"crypto/sha1"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
)

//Commit creates the commit.
func Commit(message string, files []string) (string, error) {
	treeHash, err := setFiles(files)
	if err != nil {
		return "", err
	}
	info, infoSum, err := createCommitInfo(treeHash, message)
	if err != nil {
		return "", err
	}
	err = createCommitFiles(info, infoSum)
	if err != nil {
		return "", err
	}
	return infoSum, err
}

func createCommitFiles(info string, hash string) error {
	infoFile, err := os.Create(path.Join(".svcs/commits", hash+".txt"))
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString(info)
	return err
}

func createCommitInfo(treeHash string, message string) (string, string, error) {
	head, err := GetHead()
	if err != nil {
		return "", "", err
	}
	parentSum, _, err := ConvertToCommit(head)
	if err != nil {
		return "", "", err
	}
	currentUser, err := user.Current()
	if err != nil {
		return "", "", err
	}
	info := "author " + currentUser.Username + "\ntime " + GetTime() + "\nparent " + parentSum + "\ntree " + treeHash + "\nmessage " + Encode(message)
	hash := sha1.Sum([]byte(info))
	return info, fmt.Sprintf("%x", hash), nil
}

func setFiles(files []string) (string, error) {
	content := strings.Join(files, "\n")
	hash := sha1.Sum([]byte(content))
	hashString := fmt.Sprintf("%x", hash)
	file, err := os.Create(path.Join(".svcs/trees", hashString))
	if err != nil {
		return "", err
	}
	zippedContent := Zip([]byte(content))
	_, err = file.WriteString(zippedContent)
	return hashString, err
}
