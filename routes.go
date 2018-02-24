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
