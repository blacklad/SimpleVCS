package cmd

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/MSathieu/SimpleVCS/lib"
)

func Pull(url string) error {
	conn, err := net.Dial("udp", url)
	if err != nil {
		return err
	}
	fmt.Fprint(conn, "commits")
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\r')
	if err != nil {
		return err
	}
	commitsArr := strings.Split(message, "\n")
	fmt.Fprint(conn, "branches")
	reader = bufio.NewReader(conn)
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
	reader = bufio.NewReader(conn)
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
