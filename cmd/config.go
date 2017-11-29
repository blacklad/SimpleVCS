package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/util"
)

//GetConfig prints the value of a config
func GetConfig(key string) error {
	config, err := util.GetConfig(key)
	if err != nil {
		return err
	}
	fmt.Println(config)
	return nil
}

//SetConfig sets a config key.
func SetConfig(key string, value string) error {
	split, err := gotils.SplitFileIntoArr(".svcs/settings.txt")
	if err != nil {
		return err
	}
	for i := range split {
		mapping := strings.Split(split[i], " ")
		if mapping[0] == key {
			split[i] = mapping[0] + " " + value
		}
	}
	file, err := os.Create(".svcs/settings.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(split, "\n"))
	return err
}
