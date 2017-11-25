package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/lib"
)

var response http.ResponseWriter
var auths []lib.Auth
var public bool
var multiserver bool

//Server starts the SVCS server
func Server(publicPull bool) error {
	public = publicPull
	if !gotils.CheckIfExists(".svcs") {
		multiserver = true
	}
	var err error
	auths, err = lib.GetAuth()
	if err != nil {
		return err
	}
	http.HandleFunc("/", server)
	err = http.ListenAndServeTLS("0.0.0.0:333", "svcs.crt", "svcs.key", nil)
	return err
}
func server(responseWriter http.ResponseWriter, request *http.Request) {
	response = responseWriter
	if request.Method == "GET" && request.URL.Path == "/system" {
		switch multiserver {
		case true:
			fmt.Fprint(response, "simplevcs 1.0.0 multiserver")
		case false:
			fmt.Fprint(response, "simplevcs 1.0.0 normal")
		}
		return
	}
	var authed bool
	if request.Method == "GET" && public {
		authed = true
	}
	for _, auth := range auths {
		if auth.Username == request.Header.Get("USERNAME") {
			if auth.Password == gotils.GetChecksum(request.Header.Get("PASSWORD")) {
				authed = true
			}
		}
	}
	if !authed {
		response.WriteHeader(403)
		return
	}
	var err error
	switch multiserver {
	case false:
		switch request.Method {
		case "GET":
			switch request.URL.Path {
			case "/files":
				err = lib.PullFiles(response)
			case "/trees":
				err = lib.PullTrees(response)
			case "/commits":
				err = lib.PullCommits(response)
			case "/branches":
				err = lib.PullBranches(response)
			case "/tags":
				err = lib.PullTags(response)
			default:
				response.WriteHeader(404)
				return
			}
		case "POST":
			switch request.URL.Path {
			case "/files":
				err = lib.PushFiles(request)
			case "/trees":
				err = lib.PushTrees(request)
			case "/commits":
				err = lib.PushCommits(request)
			case "/branches":
				err = lib.PushBranches(request)
			case "/tags":
				err = lib.PushTags(request)
			default:
				response.WriteHeader(404)
				return
			}
		}
	case true:
		split := strings.Split(request.URL.Path, "/")
		if strings.ContainsAny(split[0], "/\\.") {
			response.WriteHeader(400)
			return
		}
		err = os.Chdir(split[1])
		if err != nil {
			log.Println(err)
			response.WriteHeader(500)
			return
		}
		defer os.Chdir("..")
		switch request.Method {
		case "GET":
			switch split[2] {
			case "files":
				err = lib.PullFiles(response)
			case "trees":
				err = lib.PullTrees(response)
			case "commits":
				err = lib.PullCommits(response)
			case "branches":
				err = lib.PullBranches(response)
			case "tags":
				err = lib.PullTags(response)
			default:
				response.WriteHeader(404)
				return
			}
		case "POST":
			switch split[2] {
			case "files":
				err = lib.PushFiles(request)
			case "trees":
				err = lib.PushTrees(request)
			case "commits":
				err = lib.PushCommits(request)
			case "branches":
				err = lib.PushBranches(request)
			case "tags":
				err = lib.PushTags(request)
			default:
				response.WriteHeader(404)
				return
			}
		}
	}
	if err != nil {
		log.Println(err)
		response.WriteHeader(500)
	}
}
