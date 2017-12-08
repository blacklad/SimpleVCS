package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MSathieu/Gotils"
)

var response http.ResponseWriter
var auths []auth
var public bool

//Server starts the SVCS server
func Server(publicPull bool) error {
	public = publicPull
	var err error
	auths, err = getAuth()
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
	split := strings.Split(request.URL.Path, "/")
	if strings.ContainsAny(split[0], "/\\.") {
		response.WriteHeader(400)
		return
	}
	err := os.Chdir(split[1])
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
			err = pullFiles(response)
		case "trees":
			err = pullTrees(response)
		case "commits":
			err = pullCommits(response)
		case "branches":
			err = pullBranches(response)
		case "tags":
			err = pullTags(response)
		default:
			response.WriteHeader(404)
			return
		}
	case "POST":
		switch split[2] {
		case "files":
			err = pushFiles(request)
		case "trees":
			err = pushTrees(request)
		case "commits":
			err = pushCommits(request)
		case "branches":
			err = pushBranches(request)
		case "tags":
			err = pushTags(request)
		default:
			response.WriteHeader(404)
			return
		}
	}
	if err != nil {
		log.Println(err)
		response.WriteHeader(500)
	}
}
