package cmd

import (
	"errors"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Push(url string) error {
	if !lib.VCSExists() {
		return errors.New("not initialized")
	}
	return nil
}
