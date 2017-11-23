package lib

import (
	"os"

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
