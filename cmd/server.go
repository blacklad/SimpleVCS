package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MSathieu/Gotils"

	"github.com/MSathieu/SimpleVCS/lib"
)

var response http.ResponseWriter
var auths []lib.Auth

//Server starts the SVCS server
func Server() error {
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
		fmt.Fprint(response, "simplevcs 1.0.0")
		return
	}
	authed := false
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
	if request.Method == "GET" {
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
	}
	if request.Method == "POST" {
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
	if err != nil {
		response.WriteHeader(500)
		log.Println(err)
	}
}
