package util

import "io/ioutil"

//Head is the head object.
type Head struct {
	Branch   string
	Detached bool
}

//GetHead returns the head.
func GetHead() (Head, error) {
	headBytes, err := ioutil.ReadFile(".svcs/head.txt")
	if err != nil {
		return Head{}, err
	}
	head := string(headBytes)
	headObj := Head{Detached: false}
	if head == "DETACHED" {
		headObj.Detached = true
	} else {
		headObj.Branch = head
	}
	return headObj, nil
}
