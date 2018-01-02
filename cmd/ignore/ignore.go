package ignore

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/cmd/modules"
)

var ignoreList = []string{"svcs.db",
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
	"*.pid",
	"*.pid.lock",
	"dependency-reduced-pom.xml",
	"*.key",
	"*.crt"}

//CheckIgnored checks if the file/directory must be ignored.
func CheckIgnored(fileString string) (bool, error) {
	ignoreContentBytes, err := ioutil.ReadFile(".svcsignore")
	if err == nil {
		ignoreContent := string(ignoreContentBytes)
		ignoreContent = gotils.NormaliseLineEnding(ignoreContent)
		ignoreArr := strings.Split(ignoreContent, "\n")
		for _, line := range ignoreArr {
			ignoreList = append(ignoreList, line)
		}
	}
	modules, err := modules.Get()
	if err == nil {
		for _, module := range modules {
			ignoreList = append(ignoreList, module.Name)
		}
	}
	for _, line := range ignoreList {
		match, err := path.Match(line, fileString)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}
