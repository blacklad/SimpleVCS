package lib

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/util"
)

func Commit() error {
	exists := util.VCSExists(".svcs")
	if !exists {
		return errors.New("not initialized")
	}
	filepath.Walk(".", visit)
	return nil
}
func visit(filePath string, fileInfo os.FileInfo, err error) error {
	fixedPath := strings.Replace(filePath, "\\", "/", -1)
	pathArr := strings.Split(fixedPath, "/")
	for _, pathPart := range pathArr {
		if pathPart == ".svcs" {
			return nil
		}
	}
	currentPath, _ := os.Getwd()
	copyTo := path.Join(".svcs", util.GetTime())
	os.Mkdir(copyTo, 0700)
	relativePath := strings.Replace(fixedPath, currentPath, "", 1)
	newPath := path.Join(copyTo, relativePath)
	if fileInfo.IsDir() {
		os.Mkdir(newPath, fileInfo.Mode())
	} else {
		newFile, _ := os.Create(newPath)
		oldFile, _ := os.Open(fixedPath)
		io.Copy(newFile, oldFile)
	}
	return nil
}
