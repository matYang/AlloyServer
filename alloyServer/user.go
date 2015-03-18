package main

import "time"

type User struct {
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	TimeStamp time.Time `json:"timeStamp"`
}

type Users []User
