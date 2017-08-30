package lib

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

//VCSExists checks if the .svcs directory exists.
func VCSExists() bool {
	_, err := os.Stat(".svcs")
	if err != nil {
		return false
	}
	return true
}

//GetTime returns the properly formatted date and time.
func GetTime() string {
	return time.Now().Format("20060102150405")
}

//GetParent returns the parent sha of the specified commit.
func GetParent(currentSha string) (string, error) {
	currentInfo, err := ioutil.ReadFile(path.Join(".svcs/history", currentSha+".txt"))
	if err != nil {
		return "", err
	}
	splitFile := strings.Split(string(currentInfo), "\n")
	for _, line := range splitFile {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "parent" {
			if line == "parent " {
				return "", nil
			}
			return splitLine[1], nil
		}
	}
	return "", nil
}

//Zip zips the argument and returns the zipped content.
func Zip(text []byte) string {
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	comp.Write(text)
	comp.Close()
	return compBytes.String()
}

//Unzip unzips the argument and returns the normal content.
func Unzip(text []byte) string {
	var compBytes bytes.Buffer
	compBytes.Write(text)
	comp, _ := gzip.NewReader(&compBytes)
	var outputBytes bytes.Buffer
	outputBytes.ReadFrom(comp)
	comp.Close()
	return outputBytes.String()
}
