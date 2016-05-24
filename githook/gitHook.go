package gitHook

import (
   	"net/http"
	"os/exec"
	"strings"
	"os"
	"github.com/techno-chimp-www/log"
)

// some notes on what happens.
// 1 - when a push is made to the branch, github does a POST request to technochimp.com/githook/
// 2 - that triggers a bunch of processes that eventually result in a new server going live
// Processes are: git pull, go install and server die
// Restarting of the server is managed by the crontab of www-data

func Handler(res http.ResponseWriter, req *http.Request) { 
	if req.URL.Path == "/githook/" {

		log.Init(os.Stdout, os.Stdout, os.Stderr, "/var/www/sites/technochimp.com/logs/githook.log")
		log.Info.Println("githook requested");

		// STEP #1 - git pull
		gitPull()
		// STEP #2 - go install
		goInstall()
		log.Info.Printf("Turning myself off ... \n")
		// TODO: this way of turning off a server is pretty ugly, but it works like a charm!
		os.Exit(0)

		return
	}
	return
} 

func logCommand(cmd *exec.Cmd) {

	log.Info.Printf("Executing: %s\n", strings.Join(cmd.Args, " "))
}

func logOutput(outs []byte) {
 	if len(outs) > 0 {
    		log.Info.Printf("Output: %s\n", string(outs))
  	}
}

func logError(err error) {
	if err != nil {
    		log.Error.Printf("Error: %s\n", err.Error())
  	}
}	

// STEP #1 - git pull
func gitPull() {

	cmd := exec.Command("git", "-C", "/home/www-data/go/src/github.com/techno-chimp-www/", "pull")
	// here we need to be specific where to pull from -- full path to destination is appropriate
	logCommand(cmd)
	output, err := cmd.CombinedOutput()
	logError(err)
	logOutput(output)
}

// STEP #2 - go install
func goInstall() {
	cmd := exec.Command("go", "install", "github.com/techno-chimp-www")
	logCommand(cmd)
	output, err := cmd.CombinedOutput()
	logError(err)
	logOutput(output)
}

