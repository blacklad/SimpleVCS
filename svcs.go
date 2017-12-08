package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/MSathieu/Gotils"
	"github.com/MSathieu/SimpleVCS/cmd/branch"
	"github.com/MSathieu/SimpleVCS/cmd/checkout"
	"github.com/MSathieu/SimpleVCS/cmd/commit"
	"github.com/MSathieu/SimpleVCS/cmd/config"
	"github.com/MSathieu/SimpleVCS/cmd/diff"
	"github.com/MSathieu/SimpleVCS/cmd/gc"
	"github.com/MSathieu/SimpleVCS/cmd/initialize"
	"github.com/MSathieu/SimpleVCS/cmd/log"
	"github.com/MSathieu/SimpleVCS/cmd/merge"
	"github.com/MSathieu/SimpleVCS/cmd/pull"
	"github.com/MSathieu/SimpleVCS/cmd/push"
	"github.com/MSathieu/SimpleVCS/cmd/server"
	"github.com/MSathieu/SimpleVCS/cmd/stash"
	"github.com/MSathieu/SimpleVCS/cmd/stats"
	"github.com/MSathieu/SimpleVCS/cmd/status"
	"github.com/MSathieu/SimpleVCS/cmd/tag"
	"github.com/MSathieu/SimpleVCS/types"
)

func main() {
	flag.Usage = usage
	var branchVar, username, password string
	var noHead, zip, bare, public bool
	flag.StringVar(&branchVar, "branch", "master", "Specify the branch.")
	flag.BoolVar(&zip, "zip", true, "Specify if you want to zip everything when creating a project.")
	flag.BoolVar(&noHead, "no-head", false, "Don't move head.")
	flag.BoolVar(&bare, "bare", false, "Create a bare repository")
	flag.StringVar(&username, "username", "", "The username for pulling/pushing")
	flag.BoolVar(&public, "public-pull", false, "Make pulling from server public")
	flag.StringVar(&password, "password", "", "The password for pulling/pushing")
	flag.Parse()
	executedCommand := flag.Arg(0)
	if executedCommand == "" {
		flag.Usage()
		return
	}
	if executedCommand != "init" && executedCommand != "server" && !gotils.CheckIfExists(".svcs") {
		log.Fatal("not initialized")
	}
	var err error
	switch executedCommand {
	case "init":
		if gotils.CheckIfExists(".svcs") {
			log.Fatal(errors.New("already initialized"))
		}
		err = initialize.Initialize(flag.Arg(1), zip, bare)
	case "commit":
		err = commit.Commit(flag.Arg(1))
	case "config":
		if flag.Arg(2) == "" {
			err = config.GetConfig(flag.Arg(1))
		} else {
			err = config.SetConfig(flag.Arg(1), flag.Arg(2))
		}
	case "checkout":
		err = checkout.Checkout(flag.Arg(1), noHead)
	case "log":
		err = vcslog.Log(branchVar)
	case "stash":
		switch flag.Arg(1) {
		case "create":
			err = stash.CreateStash(flag.Arg(2))
		case "checkout":
			err = stash.CheckoutStash(flag.Arg(2))
		case "rm":
			err = stash.RemoveStash(flag.Arg(2))
		default:
			fmt.Println("Invalid command, run --help to get a list of the commands.")
		}
	case "patch":
		switch flag.Arg(1) {
		case "apply":
			err = types.ApplyPatch(flag.Arg(2))
		case "gen":
			err = types.GeneratePatch(flag.Arg(2), flag.Arg(3), flag.Arg(4))
		default:
			fmt.Println("Invalid command, run --help to get a list of the commands.")
		}
	case "tag":
		switch flag.Arg(1) {
		case "create":
			err = tag.Create(flag.Arg(2), flag.Arg(3))
		case "delete":
			err = tag.Remove(flag.Arg(2))
		case "list":
			err = tag.List()
		default:
			fmt.Println("Invalid command, run --help to get a list of the commands.")
		}
	case "branch":
		switch flag.Arg(1) {
		case "create":
			err = branch.CreateBranch(flag.Arg(2))
		case "delete":
			err = branch.RemoveBranch(flag.Arg(2))
		case "list":
			err = branch.ListBranches()
		default:
			fmt.Println("Invalid command, run --help to get a list of the commands.")
		}
	case "merge":
		err = merge.Merge(flag.Arg(1))
	case "pull":
		err = pull.Pull(flag.Arg(1), username, password)
	case "push":
		err = push.Push(flag.Arg(1), username, password)
	case "status":
		err = status.Status()
	case "diff":
		err = diff.Diff(flag.Arg(1), flag.Arg(2))
	case "stats":
		err = stats.ShowStats()
	case "gc":
		err = gc.GC()
	case "server":
		err = server.Server(public)
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
	fmt.Println("\ttag: The tag commands:")
	fmt.Println("\t\tlist: Lists all tags.")
	fmt.Println("\t\tcreate: Creates a tag.")
	fmt.Println("\t\tdelete: Deletes a tag.")
	fmt.Println("\tbranch: The branch commands:")
	fmt.Println("\t\tlist: Lists all branches.")
	fmt.Println("\t\tcreate: Creates a branch from the current branch.")
	fmt.Println("\t\tdelete: Deletes a branch.")
	fmt.Println("\tpull: Pulls the latest changes from the server.")
	fmt.Println("\tpush: Pushes the changes to the server.")
	fmt.Println("\tmerge: Merges two branches.")
	fmt.Println("Arguments:")
	flag.PrintDefaults()
}
