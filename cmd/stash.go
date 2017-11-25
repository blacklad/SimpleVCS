package cmd

import "os"

//CreateStash creates a stash
func CreateStash(name string) error {
	return nil
}

//CheckoutStash checkouts a stash
func CheckoutStash(name string) error {
	return nil
}

//RemoveStash removes a stash
func RemoveStash(name string) error {
	return os.Remove(".svcs/stashes/" + name)
}
