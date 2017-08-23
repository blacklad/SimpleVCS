package main

import (
	"flag"
	"os"
	"path"
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
}
