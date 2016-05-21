package log

import (
	"log"
	"io"
	"os"
	"net/http"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init( infoHandle, warningHandle, errorHandle io.Writer, filepath string) {

	Info = log.New(infoHandle,
		"INFO: \t\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: \t",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: \t\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	// set output for these 3 io.writers to go to a file
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file")
	}

	Info.SetOutput(file)
	Warning.SetOutput(file)
	Error.SetOutput(file)

}

// the following function can be user directly in the http.ListenAndServe()
func ServerLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Info.Printf("%s : %s : %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

