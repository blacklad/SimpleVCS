package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MSathieu/SimpleVCS/cmd"
)

func main() {
	var branch string
	flag.StringVar(&branch, "branch", "master", "Specify the branch")
	flag.Parse()
	executedCommand := flag.Arg(0)
	var err error
	switch executedCommand {
	case "init":
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
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
