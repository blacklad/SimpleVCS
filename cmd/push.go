package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

var body string

//Push pushes the changes to the server.
func Push(url string) error {
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
	err = filepath.Walk(".svcs/files", visitFilesPush)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/files", "", strings.NewReader(body))
	if err != nil {
		return err
	}
	return nil
}
func visitFilesPush(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	unzipped, err := lib.Unzip(string(content))
	body = body + info.Name() + " " + lib.Encode(unzipped) + "\n"
	return nil
}
