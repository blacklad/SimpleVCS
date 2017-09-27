package lib

import (
	"io/ioutil"
	"os"
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
	err = CheckIntegrity(file, hash)
	return File{Content: file, Hash: hash}, err
}

//Save saves the file
func (file File) Save() error {
	path := path.Join(".svcs/files", file.Hash)
	zippedContent := Zip(file.Content)
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = newFile.WriteString(zippedContent)
	return err
}
