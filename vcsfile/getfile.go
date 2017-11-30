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
	file, err := util.Unzip(string(zippedFile))
	if err != nil {
		return File{}, err
	}
	err = gotils.CheckIntegrity(file, hash)
	return File{Content: file, Hash: hash}, err
}
