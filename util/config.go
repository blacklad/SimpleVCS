package util

import (
	"strings"

	"github.com/MSathieu/Gotils"
)

//GetConfig gets the config.
func GetConfig(key string) (string, error) {
	split, err := gotils.SplitFileIntoArr(".svcs/settings.txt")
	if err != nil {
		return "", err
	}
	for _, line := range split {
		mapping := strings.Split(line, " ")
		if mapping[0] == key {
			return mapping[1], nil
		}
	}
	return "", nil
}

//Initialized checks whether the repo is initialized
func Initialized() bool {
	config := &Config{}
	DB.Where(&Config{Name: "name"}).First(config)
	if config.Value == "" {
		return false
	}
	return true
}
