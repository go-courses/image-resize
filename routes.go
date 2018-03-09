package main

import "net/http"

// Route there we define it
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes this is a struct
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Images",
		"GET",
		"/images",
		Images,
	},
	Route{
		"Images",
		"POST",
		"/images",
		UploadImage,
	},
	Route{
		"imageShow",
		"GET",
		"/images/{imageId}",
		ResizeImage,
	},
}
