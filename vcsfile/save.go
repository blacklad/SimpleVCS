package vcsfile

import (
	"os"
	"path"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
)

//Save saves the file
func (file *File) Save() error {
	file.Content = gotils.NormaliseLineEnding(file.Content)
	if !strings.HasSuffix(file.Content, "\n") && file.Content != "" {
		file.Content = file.Content + "\n"
	}
	file.Hash = gotils.GetChecksum(file.Content)
	path := path.Join(".svcs/files", file.Hash)
	zippedContent, err := util.Zip(file.Content)
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
