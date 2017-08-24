package lib

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"
)

func CreateMessage(time string, branch string) (string, string) {
	currentUser, _ := user.Current()
	branches, _ := ioutil.ReadFile(".svcs/branches.txt")
	var parentSum string
	branchesArr := strings.Split(string(branches), "\n")
	for _, line := range branchesArr {
		if line == "" {
			continue
		}
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == branch {
			parentSum = lineSplit[1]
			break
		}
	}
	message := "author " + currentUser.Username + "\ntime " + time + "\nparent " + parentSum
	hash := sha1.Sum([]byte(message))
	return message, fmt.Sprintf("%x", hash)
}
