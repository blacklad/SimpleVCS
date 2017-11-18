package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MSathieu/SimpleVCS/lib"
)

var response http.ResponseWriter

//Server starts the SVCS server
func Server(password string) error {
	http.HandleFunc("/", server)
	err := http.ListenAndServe("0.0.0.0:333", nil)
	return err
}
func server(responseWriter http.ResponseWriter, request *http.Request) {
	response = responseWriter
	var err error
	if request.Method == "GET" {
		switch request.URL.Path {
		case "/system":
			fmt.Fprint(response, "simplevcs 1.0.0")
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
		}
	}
	if err != nil {
		response.WriteHeader(500)
		log.Println(err)
	}
}
