package lib

import (
	"io/ioutil"
	"strings"
)

//GetConfig gets the config.
func GetConfig(key string) (string, error) {
	file, err := ioutil.ReadFile(".svcs/settings.txt")
	if err != nil {
		return "", err
	}
	split := strings.Split(string(file), "\n")
	for _, line := range split {
		mapping := strings.Split(line, " ")
		if mapping[0] == key {
			return mapping[1], nil
		}
	}
	return "", nil
}
