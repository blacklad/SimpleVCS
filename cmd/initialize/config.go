package initialize

import (
	"os"
)

func initConfig(repoName string) error {
	settingsFile, err := os.Create(".svcs/settings.txt")
	if err != nil {
		return err
	}
	_, err = settingsFile.WriteString("name " + repoName + "\n")
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
