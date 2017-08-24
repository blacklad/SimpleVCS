package lib

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
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
func GetParent(currentSha string) string {
	currentInfo, _ := ioutil.ReadFile(path.Join(".svcs/history", currentSha+".txt"))
	splitFile := strings.Split(string(currentInfo), "\n")
	for _, line := range splitFile {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "parent" {
			if line == "parent " {
				return ""
			}
			return splitLine[1]
		}
	}
	return ""
}
