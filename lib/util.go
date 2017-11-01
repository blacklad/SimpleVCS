package lib

import (
	"os"
	"strings"

	"github.com/MSathieu/Gotils"
)

//VCSExists checks if the .svcs directory exists.
func VCSExists() bool {
	_, err := os.Stat(".svcs")
	if err != nil {
		return false
	}
	return true
}

//Zip zips the argument and returns the zipped content.
func Zip(text string) (string, error) {
	config, err := GetConfig("zip")
	if err != nil {
		return "", err
	}
	if config == "false" {
		return text, nil
	}
	return gotils.GZip(text), nil
}

//Unzip unzips the argument and returns the normal content.
func Unzip(text string) (string, error) {
	config, err := GetConfig("zip")
	if err != nil {
		return "", err
	}
	if config == "false" {
		return text, nil
	}
	return gotils.UnGZip(text), nil
}

//GenerateChange returns the changes between two file arrays.
func GenerateChange(fromFiles []string, toFiles []string) []string {
	var changes []string
	for _, line := range fromFiles {
		change := "deleted"
		changed := true
		split := strings.Split(line, " ")
		for _, toLine := range toFiles {
			toSplit := strings.Split(toLine, " ")
			if split[0] == toSplit[0] {
				if split[1] == toSplit[1] {
					changed = false
				} else {
					change = "changed"
				}
			}
		}
		if changed {
			changes = append(changes, change+" "+split[0])
		}
	}
	for _, toLine := range toFiles {
		change := "created"
		changed := true
		toSplit := strings.Split(toLine, " ")
		for _, line := range fromFiles {
			split := strings.Split(line, " ")
			if split[0] == toSplit[0] {
				changed = false
			}
		}
		if changed {
			changes = append(changes, change+" "+toSplit[0])
		}
	}
	return changes
}
