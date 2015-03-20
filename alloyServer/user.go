package main

import "time"

type User struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Users []User
