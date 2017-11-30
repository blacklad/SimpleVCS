package initialize

import (
	"os"
	"strconv"
)

func initConfig(repoName string, zipped bool, bare bool, dirName string) error {
	settingsFile, err := os.Create(dirName + "/settings.txt")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("name " + repoName + "\n")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("bare " + strconv.FormatBool(bare) + "\n")
	_, err = settingsFile.WriteString("zip " + strconv.FormatBool(zipped) + "\n")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("remote " + "\n")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("username " + "\n")
	return err
}
