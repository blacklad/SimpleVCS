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
		err = cmd.Commit(branch, flag.Arg(1))
	case "checkout":
		err = cmd.Checkout(flag.Arg(1))
	case "log":
		err = cmd.Log(branch)
	case "branch":
		err = cmd.CreateBranch(flag.Arg(1), flag.Arg(2))
	case "tag":
		err = cmd.CreateTag(flag.Arg(1), flag.Arg(2))
	case "tags":
		err = cmd.ListTags()
	case "branches":
		err = cmd.ListBranches()
	case "merge":
		err = cmd.Merge(flag.Arg(1), flag.Arg(2))
	case "rmbranch":
		err = cmd.RemoveBranch(flag.Arg(1))
	case "rmtag":
		err = cmd.RemoveTag(flag.Arg(1))
	case "pull":
		err = cmd.Pull(flag.Arg(1))
	case "push":
		err = cmd.Push(flag.Arg(1))
	default:
		flag.PrintDefaults()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("Commands:")
	fmt.Println("\tCommit: Commits the current workspace to the directory specified by the branch option.")
	fmt.Println("\tcheckout: Checks out the provided sha.")
	fmt.Println("\tlog: Logs all commits of the branch specified by the branch option.")
	fmt.Println("\ttag: Creates a tag with the specified name and sha.")
	fmt.Println("\tbranch: Creates a branch with the specified name and sha.")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
}
