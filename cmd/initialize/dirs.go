package initialize

import "os"

func initDirs(dirName string) error {
	err := os.Mkdir(dirName, 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/files", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/commits", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/trees", 0700)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirName+"/stashes", 0700)
	return err
}
