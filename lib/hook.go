package lib

import gake "github.com/MSathieu/Gake/lib"

//ExecHook executes a hook.
func ExecHook(name string) error {
	gake.Process(name, true, false)
	return nil
}
