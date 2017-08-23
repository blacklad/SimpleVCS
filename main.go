package main

import (
	"flag"
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
	if executedCommand == "init" {
		lib.InitRepo(storageDir)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
