package main

import (
	"github.com/techno-chimp-www/githook"
	"github.com/techno-chimp-www/log" // wrapper around the golang's native log
	"github.com/techno-chimp-www/main-component"
	"net/http"
	"os"
)

func main() {

	log.Init(os.Stdout, os.Stdout, os.Stderr, "server.log")

	http.HandleFunc("/", mainComp.Handler)
	http.HandleFunc("/githook/", gitHook.Handler)

	http.ListenAndServe(":8000", log.ServerLog(http.DefaultServeMux))

}
