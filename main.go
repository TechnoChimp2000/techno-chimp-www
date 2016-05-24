package main

import (
	"github.com/techno-chimp-www/githook"
	"github.com/techno-chimp-www/log" // wrapper around the golang's native log
	"github.com/techno-chimp-www/main-component"
	"net/http"
	"os"
	"syscall"
	"fmt"
	"flag"
)

func main() {
	// command arguments processing
	debug := flag.Bool("debug", false, "debug option uses a default home directory.")
	flag.Parse()

	if *debug {
		fmt.Println("Debug mode enabled, special rules apply. ")

		wdBefore, _ := syscall.Getwd()
		fmt.Println("initial working directory: ", wdBefore)

		// change working directory
		syscall.Chdir("/Users/ivalantic/work/go/src/github.com/techno-chimp-www")

		wdAfter, _ := syscall.Getwd()
		fmt.Println("working directory changed to ", wdAfter )

		log.Init(os.Stdout, os.Stdout, os.Stderr, "server.log")

	} else {
		syscall.Chdir("/var/www/sites/technochimp.com/public")
		log.Init(os.Stdout, os.Stdout, os.Stderr, "/var/www/sites/technochimp.com/logs/server.log")
	}

	http.HandleFunc("/", mainComp.Handler)
	http.HandleFunc("/githook/", gitHook.Handler)

	http.ListenAndServe(":8000", log.ServerLog(http.DefaultServeMux))
}

