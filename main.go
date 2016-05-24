package main

import (
	"github.com/techno-chimp-www/githook"
	"github.com/techno-chimp-www/log" // wrapper around the golang's native log
	"github.com/techno-chimp-www/main-component"
	"net/http"
	"os"
	"syscall"
	"fmt"
)

func main() {

	var debug bool = true

	if debug {
		wdBefore, _ := syscall.Getwd()
		fmt.Println("current working directory: ", wdBefore)

		// change working directory
		syscall.Chdir("/Users/ivalantic/work/go/src/github.com/techno-chimp-www")

		wdAfter, _ := syscall.Getwd()
		fmt.Println("WD changed to: ", wdAfter )

	} else {
		syscall.Chdir("/var/www/sites/technochimp.com/public")
	}

	log.Init(os.Stdout, os.Stdout, os.Stderr, "server.log")

	http.HandleFunc("/", mainComp.Handler)
	http.HandleFunc("/githook/", gitHook.Handler)

	http.ListenAndServe(":8000", log.ServerLog(http.DefaultServeMux))

}

