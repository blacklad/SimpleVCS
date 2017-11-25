package util

import (
	"os"
	"path/filepath"

	"github.com/MSathieu/Gotils"
)

var objects []string

//GetAllObjects gets all stored objects by type
func GetAllObjects(dir string) ([]string, error) {
	err := filepath.Walk(".svcs/"+dir, visitObjects)
	objs := objects
	objects = nil
	return objs, err
}
func visitObjects(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	objects = append(objects, info.Name())
	return nil
}

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
