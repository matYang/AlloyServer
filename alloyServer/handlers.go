package main

import (
	"encoding/json"
	"github.com/matYang/AlloyServer/alsParser"
	"github.com/matYang/AlloyServer/dataModel"
	"github.com/matYang/AlloyServer/utility"
	"io"
	"io/ioutil"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user dataModel.User
	var err error
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	data := string(body[:])

	//fill in the data for user
	responseChan := make(chan dataModel.Response)
	user.Data = data
	user.Id, err = utility.UUID()
	if err != nil {
		panic(err)
	}
	user.SenderChan = &responseChan

	//request parse and await response
	alsParser.RequestParsing(user)
	//no need to manually close the channel
	response := <-responseChan

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func TestAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var response dataModel.Response
	response.Result = "Test Success"
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
