package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"io/ioutil"
	"os"
	"utility"
)

var (
	parseChan chan User
)

const (
	//TODO Find a way to initialize current directory
	CURRENTPATH = "."
	JSONPATH    = os.PathSeparator + "json"
	PYPATH      = os.PathSeparator + "py"
	ALSPATH     = os.PathSeparator + "als"
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
func PendingParsing() {
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
	err = ioutil.WriteFile(user.Name+".json", user.Data, 0644)
	if err != nil {
		panic(err)
	}

	//execute the python script, at a specific location

}
