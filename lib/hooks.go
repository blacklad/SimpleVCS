package lib

import (
	"os"

	gake "github.com/MSathieu/Gake/lib"
)

//ExecTarget executes a Gake task.
func ExecTarget(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return nil
	}
	gake.Process(path, true, false)
	return nil
}

//ExecHook provides a wrapper above ExecTarget.
func ExecHook(name string) error {
	path := ".svcs/hooks/" + name
	err := ExecTarget(path)
	return err
}
