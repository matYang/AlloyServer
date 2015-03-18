package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

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

	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	user, _ = RepoCreate(user)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
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

	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	if user, err = RepoUpdate(user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusOK)
		panic(err)
	}
}

func Read(w http.ResponseWriter, r *http.Request) {
	var user User
	var err error
	vars := mux.Vars(r)
	name := vars["name"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if user, err = RepoRead(name); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusOK)
		panic(err)
	}
}
