package lib

import (
	"crypto/sha1"
	"fmt"
	"os"
	"os/user"
	"path"
)

//CreateCommit creates and writes to the message file, info file and creates the files file.
func CreateCommit(info string, hash string, message string) error {
	infoFile, err := os.Create(path.Join(".svcs/history", hash+".txt"))
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString(info)
	if err != nil {
		return err
	}
	fileEntriesPath := path.Join(".svcs/history", hash+"_files.txt")
	_, err = os.Create(fileEntriesPath)
	if err != nil {
		return err
	}
	messageFile, err := os.Create(path.Join(".svcs/history", hash+"_message.txt"))
	if err != nil {
		return err
	}
	_, err = messageFile.WriteString(message)
	return err
}

//CreateCommitInfo returns commit info and hash.
func CreateCommitInfo(time string, branch string) (string, string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", "", err
	}
	parentSum, _, err := ConvertToCommit(branch)
	if err != nil {
		return "", "", err
	}
	message := "author " + currentUser.Username + "\ntime " + time + "\nparent " + parentSum
	hash := sha1.Sum([]byte(message))
	return message, fmt.Sprintf("%x", hash), nil
}
