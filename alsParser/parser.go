package alsParser

import (
	"fmt"
	"github.com/matYang/AlloyServer/dataModel"
	"github.com/matYang/AlloyServer/utility"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	parseChan chan dataModel.User
)

const (
	//TODO Find a way to initialize current directory
	CURRENTPATH = "."
	ALSPATH     = CURRENTPATH + string(os.PathSeparator) + "als" + string(os.PathSeparator)

	WORKLOAD = 1
)

func init() {
	parseChan = make(chan dataModel.User)
	utility.CreateDirectoryIfNotExist(ALSPATH)
}

func RequestParsing(user dataModel.User) {
	parseChan <- user
}

//Another goroutine, easier to do ATC later on
func RunParser() {
	for i := 0; i < WORKLOAD; i++ {
		go dispatch()
	}

}

func dispatch() {
	for {
		select {
		case user := <-parseChan:
			fmt.Println("Received Parsing Request")
			parseToAls(user)
		}
	}
}

func parseToAls(user dataModel.User) {
	// write whole the body
	err := ioutil.WriteFile("transcript.json", []byte(user.Data), 0644)
	if err != nil {
		panic(err)
	}

	invokeAls(user)
}

func invokeAls(user dataModel.User) {
	var response dataModel.Response
	cmd := exec.Command("sh", "test.sh")
	err := cmd.Run()
	if err != nil {
		response.Result = "Damn you Golson!! Parser not working"
		*(user.SenderChan) <- response
		return
	}

	b, err := ioutil.ReadFile("output")
	if err != nil {
		response.Result = "Alloy output not found"
	} else {
		response.Result = string(b[:])
	}
	*(user.SenderChan) <- response
}
