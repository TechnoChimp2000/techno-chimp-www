package mainComp

import(
   "io"
   "io/ioutil"
   "net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		// path relative to main.go
		indexPage, _ := ioutil.ReadFile( "main-component/index.html")
		io.WriteString(res, string(indexPage))
      return
	}
   http.ServeFile(res, req, req.URL.Path[1:])
}

