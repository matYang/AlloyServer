package main

type User struct {
	Id 	 		string
	Data 		string
	SenderChan *chan Response
}
