package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"github.com/techno-chimp-www/canvas-component"
)

func homeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		io.WriteString(res, "404, not found")
	} else {
		// assumes index.html exists
		indexPage, _ := ioutil.ReadFile("index.html")
		io.WriteString(res, string(indexPage))
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/canvas-component/", canvas.CanvasHandler)
	http.ListenAndServe(":8000", nil)
}
