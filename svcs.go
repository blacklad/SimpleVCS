package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MSathieu/SimpleVCS/checkout"
	"github.com/MSathieu/SimpleVCS/commit"
	"github.com/MSathieu/SimpleVCS/initrepo"
)

func main() {
	flag.Parse()
	executedCommand := flag.Arg(0)
	var err error
	if executedCommand == "init" {
		err = initrepo.InitRepo(flag.Arg(1))
	} else if executedCommand == "commit" {
		err = commit.Commit()
	} else if executedCommand == "checkout" {
		err = checkout.Checkout(flag.Arg(1))
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
