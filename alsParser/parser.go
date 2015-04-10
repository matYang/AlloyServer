package alsParser

import (
	"fmt"
	"github.com/matYang/AlloyServer/dataModel"
	"github.com/matYang/AlloyServer/utility"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var (
	parseChan chan dataModel.User
)

const (
	CURRENTPATH = "."
	SEPARATOR   = string(os.PathSeparator)
	ALSPATH     = CURRENTPATH + SEPARATOR + "als"

	WORKLOAD = 3
)

func init() {
	parseChan = make(chan dataModel.User)
	for i := 0; i < WORKLOAD; i++ {
		utility.CreateDirectoryIfNotExist(ALSPATH + strconv.Itoa(i))
	}
}

func RequestParsing(user dataModel.User) {
	parseChan <- user
}

//Another goroutine, easier to do ATC later on
func RunParser() {
	for i := 0; i < WORKLOAD; i++ {
		go dispatch(i)
	}

}

func dispatch(i int) {
	workerPath := ALSPATH + strconv.Itoa(i) + SEPARATOR
	for {
		select {
		case user := <-parseChan:
			fmt.Println("Received Parsing Request")
			parseToAls(user, workerPath)
		}
	}
}

func parseToAls(user dataModel.User, workerPath string) {
	// write whole the body
	err := ioutil.WriteFile(workerPath+"transcript.json", []byte(user.Data), 0644)
	if err != nil {
		panic(err)
	}

	invokeAls(user, workerPath)
}

func invokeAls(user dataModel.User, workerPath string) {
	var response dataModel.Response
	cmd := exec.Command("bash", workerPath+"solve.sh")
	err := cmd.Run()
	if err != nil {
		panic(err)
		response.Result = "Damn you Golson!! Parser not working"
		*(user.SenderChan) <- response
		return
	}

	b, err := ioutil.ReadFile(workerPath + "output")
	if err != nil {
		response.Result = "Alloy output not found"
	} else {
		response.Result = string(b[:])
	}
	*(user.SenderChan) <- response
}
