package cmd

import (
	"github.com/MSathieu/SimpleVCS/lib"
)

//Checkout checks out the specified commit.
func Checkout(commitHash string, noHead bool) error {
	return lib.Checkout(commitHash, noHead)
}
