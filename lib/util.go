package lib

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"
)

//VCSExists checks if the .svcs directory exists.
func VCSExists() bool {
	_, err := os.Stat(".svcs")
	if err != nil {
		return false
	}
	return true
}

//GetTime returns the properly formatted date and time.
func GetTime() string {
	return time.Now().Format("20060102150405")
}

//Zip zips the argument and returns the zipped content.
func Zip(text string) (string, error) {
	config, err := GetConfig("zip")
	if err != nil {
		return "", err
	}
	if config == "false" {
		return text, nil
	}
	var compBytes bytes.Buffer
	comp := gzip.NewWriter(&compBytes)
	_, err = comp.Write([]byte(text))
	if err != nil {
		return "", err
	}
	err = comp.Close()
	if err != nil {
		return "", err
	}
	return compBytes.String(), nil
}

//Unzip unzips the argument and returns the normal content.
func Unzip(text string) (string, error) {
	config, err := GetConfig("zip")
	if err != nil {
		return "", err
	}
	if config == "false" {
		return text, nil
	}
	var compBytes bytes.Buffer
	_, err = compBytes.Write([]byte(text))
	if err != nil {
		return "", err
	}
	comp, err := gzip.NewReader(&compBytes)
	if err != nil {
		return "", err
	}
	var outputBytes bytes.Buffer
	_, err = outputBytes.ReadFrom(comp)
	if err != nil {
		return "", err
	}
	err = comp.Close()
	return outputBytes.String(), err
}

//Encode base64 encodes the string.
func Encode(decoded string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(decoded))
	return encoded
}

//Decode decodes the string.
func Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}

//CheckIntegrity checks the integrity.
func CheckIntegrity(content string, hash string) error {
	if hash != GetChecksum(content) {
		return errors.New("data has been tampered with")
	}
	return nil
}

//GetChecksum gets the checksum.
func GetChecksum(data string) string {
	checksum := sha1.Sum([]byte(data))
	return fmt.Sprintf("%x", checksum)
}
