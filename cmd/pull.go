package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Pull pulls the latest changes.
func Pull(url string) error {
	systemResponse, err := http.Get(url + "/system")
	if err != nil {
		return err
	}
	systemBytes, err := ioutil.ReadAll(systemResponse.Body)
	if err != nil {
		return err
	}
	systemSplit := strings.Split(string(systemBytes), " ")
	if systemSplit[0] != "simplevcs" {
		return errors.New("unknown server")
	}
	filesResponse, err := http.Get(url + "/files")
	if err != nil {
		return err
	}
	filesBytes, err := ioutil.ReadAll(filesResponse.Body)
	if err != nil {
		return err
	}
	filesSplit := strings.Split(string(filesBytes), "\n")
	for _, file := range filesSplit {
		fileSplit := strings.Split(file, " ")
		_, err := lib.GetFile(fileSplit[0])
		if err != nil {
			continue
		}
		decodedFile, err := lib.Decode(fileSplit[1])
		if err != nil {
			return err
		}
		fileObj := lib.File{Hash: fileSplit[0], Content: decodedFile}
		err = fileObj.Save()
		if err != nil {
			return err
		}
	}
	return nil
}
