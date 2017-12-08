package types

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
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
	fileContent := gotils.UnGZip(string(zippedFile))
	err = gotils.CheckIntegrity(fileContent, hash)
	return File{Content: fileContent, Hash: hash}, err
}

//Save saves the file
func (fileObj *File) Save() error {
	fileObj.Content = gotils.NormaliseLineEnding(fileObj.Content)
	if !strings.HasSuffix(fileObj.Content, "\n") && fileObj.Content != "" {
		fileObj.Content = fileObj.Content + "\n"
	}
	fileObj.Hash = gotils.GetChecksum(fileObj.Content)
	path := path.Join(".svcs/files", fileObj.Hash)
	zippedContent := gotils.GZip(fileObj.Content)
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = newFile.WriteString(zippedContent)
	return err
}
