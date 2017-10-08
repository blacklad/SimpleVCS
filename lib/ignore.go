package lib

import (
	"io/ioutil"
	"path"
	"strings"
)

//CheckIgnored checks if the file/directory must be ignored.
func CheckIgnored(file string) (bool, error) {
	switch file {
	case ".svcs", ".git", ".svn", ".hg":
		return true, nil
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
