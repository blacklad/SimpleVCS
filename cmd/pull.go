package cmd

import (
	"github.com/MSathieu/SimpleVCS/lib"
)

//Pull pulls the latest changes.
func Pull(url string, username string, password string) error {
	return lib.Pull(url, username, password)
}
