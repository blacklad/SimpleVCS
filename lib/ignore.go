package lib

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
)

var ignoreList = []string{".svcs",
	".git",
	".svn",
	".hg",
	"*.o",
	"*.exe",
	"*.log",
	"*.out",
	"*.gem",
	"*.zip",
	"*.tar",
	"*.jar",
	"*.war",
	"*.class",
	".idea",
	"*.iml",
	"*.swp",
	"*.iso",
	"*.pid"
	"*.pid.lock",
  "dependency-reduced-pom.xml"}

//CheckIgnored checks if the file/directory must be ignored.
func CheckIgnored(file string) (bool, error) {
	ignoreContentBytes, err := ioutil.ReadFile(".svcsignore")
	if err == nil {
		ignoreContent := string(ignoreContentBytes)
		ignoreContent = gotils.NormaliseLineEnding(ignoreContent)
		ignoreArr := strings.Split(ignoreContent, "\n")
		for _, line := range ignoreArr {
			ignoreList = append(ignoreList, line)
		}
	}
	for _, line := range ignoreList {
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
