package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Push pushes the changes to the server.
func Push(url string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	return nil
}
