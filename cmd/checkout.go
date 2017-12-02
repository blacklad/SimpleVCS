package cmd

import (
	"github.com/MSathieu/SimpleVCS/checkout"
	"github.com/MSathieu/SimpleVCS/lib"
)

//Checkout checks out the specified commit.
func Checkout(commitHash string, noHead bool) error {
	err := checkout.Checkout(commitHash, noHead)
	if err != nil {
		return err
	}
	return lib.InitModules()
}
