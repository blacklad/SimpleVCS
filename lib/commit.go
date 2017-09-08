package lib

import (
	"crypto/sha1"
	"fmt"
	"os"
	"os/user"
	"path"
)

//Commit creates the commit.
func Commit(message string, files []string) (string, error) {
	info, infoSum, err := getCommitInfo()
	if err != nil {
		return "", err
	}
	err = createCommitFiles(info, infoSum, message)
	if err != nil {
		return "", err
	}
	setFiles(infoSum, files)
	return infoSum, err
}

func createCommitFiles(info string, hash string, message string) error {
	infoFile, err := os.Create(path.Join(".svcs/commits", hash+".txt"))
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString(info)
	if err != nil {
		return err
	}
	messageFile, err := os.Create(path.Join(".svcs/commits", hash+"_message.txt"))
	if err != nil {
		return err
	}
	_, err = messageFile.WriteString(message)
	return err
}

func getCommitInfo() (string, string, error) {
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
	message := "author " + currentUser.Username + "\ntime " + GetTime() + "\nparent " + parentSum
	hash := sha1.Sum([]byte(message))
	return message, fmt.Sprintf("%x", hash), nil
}

func setFiles(commitHash string, files []string) error {
	file, err := os.Create(path.Join(".svcs/trees", commitHash+".txt"))
	var content string
	if err != nil {
		return err
	}
	for _, line := range files {
		content = content + line + "\n"
	}
	zippedContent := Zip([]byte(content))
	_, err = file.WriteString(zippedContent)
	return err
}
