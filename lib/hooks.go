package lib

import (
	"os"
	"os/user"

	gosh "github.com/MSathieu/Gosh/lib"
)

//ExecScript executes a script.
func ExecScript(path string) error {
	_, err := os.Stat(path + ".gosh")
	if err != nil {
		return nil
	}
	gosh.RunFile(path + ".gosh")
	return nil
}

//ExecHook provides a wrapper above ExecScript.
func ExecHook(name string) error {
	path := ".svcshooks/" + name
	err := ExecScript(path)
	if err != nil {
		return err
	}
	user, err := user.Current()
	if err != nil {
		return err
	}
	userPath := user.HomeDir + "/.svcshooks/" + name
	err = ExecScript(userPath)
	return err
}
