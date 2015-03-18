package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Create",
		"POST",
		"/api/susie/game1",
		Create,
	},
	Route{
		"Read",
		"GET",
		"/api/susie/game1/{name}",
		Read,
	},
	Route{
		"Update",
		"PUT",
		"/api/susie/game1/{name}",
		Update,
	},
}
