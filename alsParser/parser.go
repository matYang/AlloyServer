package alsParser

import (
	"fmt"
	"github.com/matYang/AlloyServer/dataModel"
	"github.com/matYang/AlloyServer/utility"
	"io/ioutil"
	"os"
)

var (
	parseChan chan dataModel.User
)

const (
	//TODO Find a way to initialize current directory
	CURRENTPATH = "."
	JSONPATH    = CURRENTPATH + string(os.PathSeparator) + "json" + string(os.PathSeparator)
	PYPATH      = CURRENTPATH + string(os.PathSeparator) + "py" + string(os.PathSeparator)
	ALSPATH     = CURRENTPATH + string(os.PathSeparator) + "als" + string(os.PathSeparator)

	WORKLOAD = 3
)

func init() {
	parseChan = make(chan dataModel.User)
	utility.CreateDirectoryIfNotExist(JSONPATH)
	utility.CreateDirectoryIfNotExist(PYPATH)
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
	err := ioutil.WriteFile(JSONPATH+user.Id+".json", []byte(user.Data), 0644)
	if err != nil {
		panic(err)
	}

	//execute the python script, at a specific location

	invokeAls(user)
}

func invokeAls(user dataModel.User) {

	//run alloy here
	returnToSender(user)
}

func returnToSender(user dataModel.User) {

	var response dataModel.Response

	//send response back using sender channel
	*(user.SenderChan) <- response
}
