package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/unrolled/render"
)

var rend = render.New()

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func Images(w http.ResponseWriter, r *http.Request) {
	rend.HTML(w, http.StatusOK, "index", "")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", "Ваш файл был успешно загружен!")
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		filepath := handler.Filename

		defer f.Close()
		io.Copy(f, file)

		Resize(filepath)
	}
}
