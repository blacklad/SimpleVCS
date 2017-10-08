package lib

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
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
	file, err := Unzip(string(zippedFile))
	if err != nil {
		return File{}, err
	}
	err = CheckIntegrity(file, hash)
	return File{Content: file, Hash: hash}, err
}

//AddFile adds a file to the database.
func AddFile(content string) (File, error) {
	content = strings.Replace(content, "\r\n", "\n", -1)
	content = strings.Replace(content, "\r", "\n", -1)
	if !strings.HasSuffix(content, "\n") {
		content = content + "\n"
	}
	file := File{Content: content, Hash: GetChecksum(content)}
	err := file.Save()
	return file, err
}

//Save saves the file
func (file File) Save() error {
	path := path.Join(".svcs/files", file.Hash)
	zippedContent, err := Zip(file.Content)
	if err != nil {
		return err
	}
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = newFile.WriteString(zippedContent)
	if err != nil {
		return err
	}
	err = newFile.Close()
	return err
}
