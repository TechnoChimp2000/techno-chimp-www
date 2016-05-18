package gitHook

import (
   	"net/http"
   	"fmt"
	"os/exec"
	"strings"
	"os"
)

// The handler is called by the 3rd party, github speficifally
// a push to the production branch will trigger the following:
// 1) pull into the local production branch on the server
// 2) compiling the server program from the local branch
// 3) moving the server program to where it belongs to, and restarting the server

func Handler(res http.ResponseWriter, req *http.Request) { 
	if req.URL.Path == "/githook/" {
		fmt.Println("githook requested");

		cmd := exec.Command("git", "status")
		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printError(err)
		printOutput(output)
		return
	}
	return
} 

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
 	if len(outs) > 0 {
    		fmt.Printf("==> Output: %s\n", string(outs))
  	}
}

func printError(err error) {
	if err != nil {
    		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  	}
}	
