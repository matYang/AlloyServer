package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"bytes"
	"utility"
	"alsParser"

	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user User
	var err error
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	arrSize := bytes.Index(body, []byte{0})
	data := string(body[:n])

	//fill in the data for user
	responseChan := make(chan Response);
	var user User
	user.Data = data
	user.Id, err = utility.UUID()
	if err != nil {
		panic (err)
	}
	user.SenderChan = &responseChan

	//request parse and await response
	alsParser.RequestParsing(user);
	//no need to manually close the channel
	response := <-responseChan

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
