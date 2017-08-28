package cmd

import (
	"bufio"
	"fmt"
	"net"
	"strings"
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
	fmt.Fprint(conn, "tags")
	reader = bufio.NewReader(conn)
	message, err = reader.ReadString('\r')
	if err != nil {
		return err
	}
	tagsArr := strings.Split(message, "\n")
	return nil
}
