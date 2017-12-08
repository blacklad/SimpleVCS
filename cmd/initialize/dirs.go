package initialize

import "os"

func initDirs() error {
	err := os.Mkdir(".svcs", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/files", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/commits", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/trees", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(".svcs/stashes", 0700)
	return err
}
