package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/MSathieu/SimpleVCS/cmd"
	"github.com/MSathieu/SimpleVCS/lib"
)

func main() {
	flag.Usage = usage
	var branch string
	flag.StringVar(&branch, "branch", "master", "Specify the branch.")
	flag.Parse()
	executedCommand := flag.Arg(0)
	var err error
	if executedCommand != "init" && !lib.VCSExists() {
		log.Fatal("not initialized")
	}
	switch executedCommand {
	case "init":
		if lib.VCSExists() {
			log.Fatal(errors.New("already initialized"))
		}
		err = cmd.InitRepo(flag.Arg(1))
	case "commit":
		err = cmd.Commit(flag.Arg(1))
	case "checkout":
		err = cmd.Checkout(flag.Arg(1))
	case "log":
		err = cmd.Log(branch)
	case "tag":
		switch flag.Arg(1) {
		case "create":
			err = cmd.CreateTag(flag.Arg(2), flag.Arg(3))
		case "delete":
			err = cmd.RemoveTag(flag.Arg(2))
		case "list":
			err = cmd.ListTags()
		default:
			fmt.Println("Invalid command, run --help to get a list of the commands.")
		}
	case "branch":
		switch flag.Arg(1) {
		case "create":
			err = cmd.CreateBranch(flag.Arg(2))
		case "delete":
			err = cmd.RemoveBranch(flag.Arg(2))
		case "list":
			err = cmd.ListBranches()
		default:
			fmt.Println("Invalid command, run --help to get a list of the commands.")
		}
	case "merge":
		err = cmd.Merge(flag.Arg(1))
	case "pull":
		err = cmd.Pull(flag.Arg(1))
	case "push":
		err = cmd.Push(flag.Arg(1))
	case "ignore":
		err = cmd.Ignore(flag.Arg(1))
	case "unignore":
		err = cmd.UnIgnore(flag.Arg(1))
	default:
		fmt.Println("Invalid command, run --help to get a list of the commands.")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("Commands:")
	fmt.Println("\tCommit: Commits the current workspace.")
	fmt.Println("\tcheckout: Checks out the provided branch/sha.")
	fmt.Println("\tlog: Logs all commits of the branch specified by the branch option.")
	fmt.Println("\ttag: The tag command.")
	fmt.Println("\tbranch: The branch command.")
	fmt.Println("\tpull: Pulls the latest changes from the server.")
	fmt.Println("\tpush: Pushes the changes to the server.")
	fmt.Println("\tmerge: Merges two branches.")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
}
