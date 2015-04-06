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
	JSONPATH    = string(os.PathSeparator) + "json" + string(os.PathSeparator)
	PYPATH      = string(os.PathSeparator) + "py" + string(os.PathSeparator)
	ALSPATH     = string(os.PathSeparator) + "als" + string(os.PathSeparator)

	WORKLOAD = 3
)

func init() {
	parseChan = make(chan dataModel.User)
	utility.CreateDirectoryIfNotExist(CURRENTPATH + JSONPATH)
	utility.CreateDirectoryIfNotExist(CURRENTPATH + PYPATH)
	utility.CreateDirectoryIfNotExist(CURRENTPATH + ALSPATH)
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

}

func returnToSender(user dataModel.User) {

	var response dataModel.Response

	//send response back using sender channel
	*(user.SenderChan) <- response
}