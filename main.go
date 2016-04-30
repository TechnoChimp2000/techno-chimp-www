package main

import (
   "net/http"
   "github.com/techno-chimp-www/main-component"
	"github.com/techno-chimp-www/canvas-component"
)

func main() {
	http.HandleFunc("/", mainComp.Handler)
	http.HandleFunc("/canvas-component/", canvasComp.Handler)
	http.ListenAndServe(":8000", nil)
}
