package util

import (
	"os"
	"path/filepath"
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
