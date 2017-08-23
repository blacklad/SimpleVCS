package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/MSathieu/SimpleVCS/lib"
)

const (
	name        = "SimpleVCS"
	commandName = "svcs"
	vcsDir      = ".svcs"
)

func main() {
	flag.Parse()
	currentDir, _ := os.Getwd()
	storageDir := path.Join(currentDir, vcsDir)
	executedCommand := flag.Arg(0)
	repoName := flag.Arg(1)
	var err error
	if executedCommand == "init" {
		err = lib.InitRepo(storageDir, repoName)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
