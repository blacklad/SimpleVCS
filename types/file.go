package types

import (
	"os"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
)

//File is the file object.
type File struct {
	Content string
	Hash    string
}

//Save saves the file
func (fileObj *File) Save() error {
	fileObj.Content = gotils.NormaliseLineEnding(fileObj.Content)
	if !strings.HasSuffix(fileObj.Content, "\n") && fileObj.Content != "" {
		fileObj.Content = fileObj.Content + "\n"
	}
	fileObj.Hash = gotils.GetChecksum(fileObj.Content)
	path := path.Join(".svcs/files", fileObj.Hash)
	zippedContent, err := util.Zip(fileObj.Content)
	if err != nil {
		return err
	}
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = newFile.WriteString(zippedContent)
	return err
}
