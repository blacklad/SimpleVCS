package lib

import (
	"os"
	"os/user"

	gosh "github.com/MSathieu/Gosh/cmd"
)

//ExecScript executes a script.
func ExecScript(path string) error {
	_, err := os.Stat(path + ".gosh")
	if err != nil {
		return nil
	}
	return gosh.RunFile(path + ".gosh")
}

//ExecHook provides a wrapper above ExecScript.
func ExecHook(name string) error {
	err := ExecScript(".svcshooks/" + name)
	if err != nil {
		return err
	}
	user, err := user.Current()
	if err != nil {
		return err
	}
	return ExecScript(user.HomeDir + "/.svcshooks/" + name)
}
