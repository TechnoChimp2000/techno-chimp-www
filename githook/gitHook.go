package gitHook

import (
   	"net/http"
	"os/exec"
	"strings"
	"os"
	"github.com/techno-chimp-www/log"
)

// The handler is called by the 3rd party, github speficifally
// a push to the production branch will trigger the following:
// 1) pull into the local production branch on the server
// 2) compiling the server program from the local branch
// 3) moving the server program to where it belongs to, and restarting the server

func Handler(res http.ResponseWriter, req *http.Request) { 
	if req.URL.Path == "/githook/" {

		log.Init(os.Stdout, os.Stdout, os.Stderr, "githook.log")

		log.Info.Println("githook requested");

		cmd := exec.Command("git", "status") // here we need to be specific where to pull from -- full path to destination is appropriate
		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printError(err)
		printOutput(output)
		return
	}
	return
} 

func printCommand(cmd *exec.Cmd) {

	log.Info.Printf("Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
 	if len(outs) > 0 {
    		log.Info.Printf("Output: %s\n", string(outs))
  	}
}

func printError(err error) {
	if err != nil {
    		log.Error.Printf("Error: %s\n", err.Error())
  	}
}	
