package main

import (
  "fmt"
  "github.com/unrolled/render"
  "net/http"
)

var rend = render.New()

func Index(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/images", 302)
}

func Images(w http.ResponseWriter, r *http.Request) {
  rend.HTML(w, http.StatusOK, "index", "")
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("I'm uploading images")
}

func ResizeImage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("We'll resize images from here!")
}
