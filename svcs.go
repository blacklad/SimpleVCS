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
	if executedCommand == "init" {
		err = cmd.InitRepo(flag.Arg(1))
	} else if executedCommand == "commit" {
		err = cmd.Commit(branch)
	} else if executedCommand == "checkout" {
		err = cmd.Checkout(flag.Arg(1))
	} else if executedCommand == "log" {
		err = cmd.Log(branch)
	} else if executedCommand == "branch" {
		err = cmd.CreateBranch(flag.Arg(1), flag.Arg(2))
	} else if executedCommand == "tag" {
		err = cmd.CreateTag(flag.Arg(1), flag.Arg(2))
	} else if executedCommand == "tags" {
		err = cmd.ListTags()
	} else if executedCommand == "branches" {
		err = cmd.ListBranches()
	} else if executedCommand == "merge" {
		err = cmd.Merge(flag.Arg(1), flag.Arg(2))
	} else if executedCommand == "rmbranch" {
		err = cmd.RemoveBranch(flag.Arg(1))
	} else if executedCommand == "rmtag" {
		err = cmd.RemoveTag(flag.Arg(1))
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
