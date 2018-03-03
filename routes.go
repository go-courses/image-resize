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
		"Show the Form",
		"GET",
		"/images",
		Images,
	},
	Route{
		"Upload Image",
		"POST",
		"/images",
		UploadImage,
	},
}
