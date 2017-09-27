package lib

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

//File is the file object.
type File struct {
	Content string
	Hash    string
}

//GetFile gets a file.
func GetFile(hash string) (File, error) {
	if hash == "" {
		return File{}, nil
	}
	zippedFile, err := ioutil.ReadFile(path.Join(".svcs/files", hash))
	if err != nil {
		return File{}, err
	}
	file := Unzip(string(zippedFile))
	newSha := sha1.Sum([]byte(file))
	if hash != fmt.Sprintf("%x", newSha) {
		return File{}, errors.New("data has been tampered with")
	}
	return File{Content: file, Hash: hash}, nil
}
