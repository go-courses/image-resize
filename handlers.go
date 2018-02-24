package main

import (
  "fmt"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Println("I'm working!")
}

func Images(w http.ResponseWriter, r *http.Request) {
  fmt.Println("We'll upload images from here!")
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("I'm uploading images")
}

func ResizeImage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("We'll resize images from here!")
}
