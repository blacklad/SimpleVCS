package cmd

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

//Pull pulls the latest changes.
func Pull(url string) error {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return err
	}
	fmt.Fprint(conn, "commits")
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\r')
	if err != nil {
		return err
	}
	//commitsArr := strings.Split(message, "\n")
	fmt.Fprint(conn, "branches")
	message, err = reader.ReadString('\r')
	if err != nil {
		return err
	}
	branchesArr := strings.Split(message, "\n")
	for _, branch := range branchesArr {
		mapping := strings.Split(branch, " ")
		lib.UpdateBranch(mapping[0], mapping[1])
	}
	fmt.Fprint(conn, "tags")
	message, err = reader.ReadString('\r')
	if err != nil {
		return err
	}
	tagsArr := strings.Split(message, "\n")
	for _, tag := range tagsArr {
		mapping := strings.Split(tag, " ")
		lib.CreateTag(mapping[0], mapping[1])
	}
	return nil
}
