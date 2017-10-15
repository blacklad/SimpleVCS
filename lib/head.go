package lib

import "io/ioutil"

//Head is the head object.
type Head struct {
	Branch   Branch
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
		branch, err := GetBranch(head)
		if err != nil {
			return Head{}, err
		}
		headObj.Branch = branch
	}
	return headObj, err
}
