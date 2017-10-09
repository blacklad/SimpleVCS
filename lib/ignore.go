package lib

import (
	"io/ioutil"
	"path"
	"strings"
)

var ignoreList = []string{".svcs", ".git", ".svn", ".hg"}

//CheckIgnored checks if the file/directory must be ignored.
func CheckIgnored(file string) (bool, error) {
	for _, value := range ignoreList {
		match, err := path.Match(value, file)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	ignoreContent, err := ioutil.ReadFile(".svcs/ignore.txt")
	if err != nil {
		return false, err
	}
	ignoreArr := strings.Split(string(ignoreContent), "\n")
	for _, line := range ignoreArr {
		match, err := path.Match(line, file)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}
