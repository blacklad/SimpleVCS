package vcsfile

import (
	"io/ioutil"
	"path"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
)

//GetFile gets a file.
func GetFile(hash string) (File, error) {
	if hash == "" {
		return File{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/files", hash))
	if err != nil {
		return File{}, err
	}
	fileContent, err := util.Unzip(string(zippedFile))
	if err != nil {
		return File{}, err
	}
	err = gotils.CheckIntegrity(fileContent, hash)
	return File{Content: fileContent, Hash: hash}, err
}
