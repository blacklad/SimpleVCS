package types

import (
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
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
	file := &util.File{}
	util.DB.Where(&File{Hash: hash}).First(file)
	err := gotils.CheckIntegrity(file.Content, file.Hash)
	return File{Content: file.Content, Hash: file.Hash}, err
}

//Save saves the file
func (fileObj *File) Save() {
	fileObj.Content = gotils.NormaliseLineEnding(fileObj.Content)
	if !strings.HasSuffix(fileObj.Content, "\n") && fileObj.Content != "" {
		fileObj.Content = fileObj.Content + "\n"
	}
	fileObj.Hash = gotils.GetChecksum(fileObj.Content)
	util.DB.Create(&util.File{Hash: fileObj.Hash, Content: fileObj.Content})
}
