package main

import (
	"os"
	"io"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/unrolled/render"
	"gopkg.in/h2non/filetype.v1"
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
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		imageStatus, err := ChekFile(handler.Filename, w)
		if err != nil {
			return
		}
		fmt.Fprintf(w, "%v %v", imageStatus, "Ваш файл прошёл проверку и успешно загружен!")

		defer f.Close()
		io.Copy(f, file)

		Resize(handler.Filename)
	}
}

func ChekFile(fileName string, w http.ResponseWriter) (bool, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false, err
	}
	kind, err := filetype.Match(file)
	if (kind.MIME.Value == "image/png" || kind.MIME.Value == "image/jpeg" || kind.MIME.Value == "image/gif") && err != nil && filetype.IsImage(file) {
		fmt.Fprintf(w, "%v", "Ваш файл это не изображение! Попробуйте снова!")
		return false, err
	}
	fmt.Printf("File type: %s. MIME: %s. IsImage: %v\n", kind.Extension, kind.MIME.Value, filetype.IsImage(file))
	return true, nil
}