package lib

import "io/ioutil"

//GetHead returns the head.
func GetHead() (string, error) {
	head, err := ioutil.ReadFile(".svcs/head.txt")
	return string(head), err
}
