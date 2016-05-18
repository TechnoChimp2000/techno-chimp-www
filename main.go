package main

import (
	"io"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!  - Igor's test go server")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println(http.ListenAndServe(":8000", nil))
}
