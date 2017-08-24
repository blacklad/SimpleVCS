package util

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"time"
)

func VCSExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return true
}
func GetTime() string {
	return time.Now().Format("20060102150405")
}
func CreateMessage(time string) (string, string) {
	currentUser, _ := user.Current()
	parentSum, _ := ioutil.ReadFile(".svcs/branches.txt")
	message := "author " + currentUser.Username + "\ntime " + time + "\nparent " + string(parentSum)
	hash := sha1.Sum([]byte(message))
	return message, fmt.Sprintf("%x", hash)
}
