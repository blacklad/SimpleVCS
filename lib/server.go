package lib

import (
	"io/ioutil"
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
	authFile, err := ioutil.ReadFile("auth.txt")
	if err != nil {
		return nil, err
	}
	auth := gotils.NormaliseLineEnding(string(authFile))
	split := strings.Split(auth, "\n")
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
