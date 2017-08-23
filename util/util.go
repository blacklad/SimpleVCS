package util

import (
	"os"
	"time"
)

func VCSExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return true
}
func GetTime() string {
	return time.Now().Format("20060102150405")
}
