package main

import (
	"errors"
)

var users Users

// Give us some seed data
func init() {
	RepoCreate(User{Name: "Write presentation"})
	RepoCreate(User{Name: "Host meetup"})
}

func RepoRead(name string) (User, error) {
	for _, t := range users {
		if t.Name == name {
			return t, nil
		}
	}
	// return empty Todo if not found
	return User{}, errors.New("NOTFOUND")
}

func RepoCreate(u User) (User, error) {
	users = append(users, u)
	return u, nil
}

func RepoUpdate(u User) (User, error) {
	for i, t := range users {
		if t.Name == u.Name {
			users[i] = u
			return u, nil
		}
	}
	// return empty Todo if not found
	return User{}, errors.New("NOTFOUND")
}
