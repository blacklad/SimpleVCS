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
	file := File{Content: content}
	err := file.Save()
	return file, err
}

//Save saves the file
func (file File) Save() error {
	file.Content = strings.Replace(file.Content, "\r\n", "\n", -1)
	file.Content = strings.Replace(file.Content, "\r", "\n", -1)
	if !strings.HasSuffix(file.Content, "\n") {
		file.Content = file.Content + "\n"
	}
	file.Hash = GetChecksum(file.Content)
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
