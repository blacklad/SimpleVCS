package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MSathieu/SimpleVCS/lib"
)

func main() {
	flag.Parse()
	executedCommand := flag.Arg(0)
	var err error
	if executedCommand == "init" {
		err = lib.InitRepo(flag.Arg(1))
	} else if executedCommand == "commit" {
		err = lib.Commit()
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
