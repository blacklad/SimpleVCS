package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MSathieu/SimpleVCS/cmd"
)

func main() {
	flag.Parse()
	executedCommand := flag.Arg(0)
	var err error
	if executedCommand == "init" {
		err = cmd.InitRepo(flag.Arg(1))
	} else if executedCommand == "commit" {
		err = cmd.Commit()
	} else if executedCommand == "checkout" {
		err = cmd.Checkout(flag.Arg(1))
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
