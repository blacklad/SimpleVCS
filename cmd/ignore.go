package cmd

import (
	"io/ioutil"
	"os"
	"strings"
)

//Ignore ignores the file.
func Ignore(fileName string) error {
	file, err := os.OpenFile(".svcs/ignore.txt", os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	_, err = file.WriteString(fileName + "\n")
	return err
}

//UnIgnore unignores the file.
func UnIgnore(fileName string) error {
	fileContent, err := ioutil.ReadFile(".svcs/ignore.txt")
	if err != nil {
		return err
	}
	fileArr := strings.Split(string(fileContent), "\n")
	file, err := os.Create(".svcs/ignore.txt")
	if err != nil {
		return err
	}
	for _, line := range fileArr {
		if line == "" {
			continue
		}
		if line == fileName {
			continue
		}
		file.WriteString(line + "\n")
	}
	return nil
}
