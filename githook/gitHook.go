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

		log.Init(os.Stdout, os.Stdout, os.Stderr, "/var/www/sites/technochimp.com/logs/githook.log")

		log.Info.Println("githook requested");

		// STEP #1 - git pull
		gitPull()
		// STEP #2 - go install
		goInstall()
		// STEP #3 - cp source target
		moveExe()
		// STEP #4 - kill the existing server if still online
		log.Info.Printf("Turning myself off ... ")
		os.Exit(0)
		// STEP #5 - restart the server ( this might be unnecessary because our cronjob will take care of it himself

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

// STEP #3 - move bin
func moveExe() {
	cmd := exec.Command("cp", "/home/www-data/go/bin/techno-chimp-www", "/var/www/sites/technochimp.com/techno-chimp-www")
	logCommand(cmd)
	output, err := cmd.CombinedOutput()
	logError(err)
	logOutput(output)
}



