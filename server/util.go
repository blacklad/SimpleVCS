package server

import (
	"strings"

	"github.com/MSathieu/Gotils"
)

type auth struct {
	Username string
	Password string
}

func getAuth() ([]auth, error) {
	split, err := gotils.SplitFileIntoArr("auth.txt")
	if err != nil {
		return nil, err
	}
	var authArr []auth
	for _, line := range split {
		if line == "" {
			continue
		}
		splitLine := strings.Fields(line)
		authArr = append(authArr, auth{Username: splitLine[0], Password: splitLine[1]})
	}
	return authArr, nil
}
