package alsParser

import (
	"fmt"
	"github.com/kardianos/osext"
	"io/ioutil"
	"os"
	"utility"
	"main"
)

var (
	parseChan chan User
)

const (
	//TODO Find a way to initialize current directory
	CURRENTPATH = "."
	JSONPATH    = os.PathSeparator + "json" + os.PathSeparator
	PYPATH      = os.PathSeparator + "py" + os.PathSeparator
	ALSPATH     = os.PathSeparator + "als" + os.PathSeparator

	WORKLOAD 	= 3;
)

func init() {
	parseChan = make(chan User)
	utility.CreateDirectoryIfNotExist(CURRENTPATH + JSONPATH)
	utility.CreateDirectoryIfNotExist(CURRENTPATH + PYPATH)
	utility.CreateDirectoryIfNotExist(CURRENTPATH + ALSPATH)
}

func RequestParsing(user User) {
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

func parseToAls(user User) {
	// write whole the body
	err = ioutil.WriteFile(JSONPATH + user.Id + ".json", user.Data, 0644)
	if err != nil {
		panic(err)
	}

	//execute the python script, at a specific location


	invokeAls(user);
}


func invokeAls(user User) {

}

func returnToSender(user User) {


	var response Response;

	//send response back using sender channel
	*(user.SenderChan) <- response;
}