package dataModel

type User struct {
	Id         string
	Data       string
	SenderChan *chan Response
}
