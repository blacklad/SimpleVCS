package lib

import (
	"strings"

	"github.com/MSathieu/Gotils"
)

//Auth is the auth object
type Auth struct {
	Username string
	Password string
}

//GetAuth reads the auth file.
func GetAuth() ([]Auth, error) {
	split, err := gotils.SplitFileIntoArr("auth.txt")
	if err != nil {
		return nil, err
	}
	var authArr []Auth
	for _, line := range split {
		if line == "" {
			continue
		}
		splitLine := strings.Fields(line)
		authArr = append(authArr, Auth{Username: splitLine[0], Password: splitLine[1]})
	}
	return authArr, nil
}
