package vcsfile

import (
	"io/ioutil"
	"path"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/types"
	"github.com/MSathieu/SimpleVCS/util"
)

//GetFile gets a file.
func GetFile(hash string) (types.File, error) {
	if hash == "" {
		return types.File{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/files", hash))
	if err != nil {
		return types.File{}, err
	}
	fileContent, err := util.Unzip(string(zippedFile))
	if err != nil {
		return types.File{}, err
	}
	err = gotils.CheckIntegrity(fileContent, hash)
	return types.File{Content: fileContent, Hash: hash}, err
}
