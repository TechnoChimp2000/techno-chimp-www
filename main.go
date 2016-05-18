package main

import (
	"net/http"
	"fmt"
	"github.com/techno-chimp-www/main-component"
	"github.com/techno-chimp-www/githook"
)

func main() {
	http.HandleFunc("/", mainComp.Handler)
	http.HandleFunc("/githook/", gitHook.Handler)
	fmt.Println(http.ListenAndServe(":8000", nil))
}
