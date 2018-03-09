package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
	"golang.org/x/image/draw"
	"gopkg.in/h2non/filetype.v1"
)

var rend = render.New()
var workingDirectory = "imgs/"

// Index root url
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/images", 302)
}

// Images page with upload form
func Images(w http.ResponseWriter, r *http.Request) {
	rend.HTML(w, http.StatusOK, "index", "")
}

// UploadImage here the form to upload img
func UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer http.Redirect(w, r, "/images/"+handler.Filename, 302)
	defer file.Close()

	if _, err := os.Stat(workingDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(workingDirectory, os.ModePerm); err != nil {
			log.Println(err)
		}
	}

	f, err := os.OpenFile(workingDirectory+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	imageStatus, err := ChekFile(workingDirectory+handler.Filename, w)
	if err != nil {
		return
	}
	fmt.Printf("%v %v", imageStatus, "Ваш файл прошёл проверку и успешно загружен!")

	defer f.Close()
	io.Copy(f, file)

}

// ChekFile for validation of image
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

// ResizeImage function to resize img
func ResizeImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var buff bytes.Buffer
	imageID := vars["imageId"]
	src, imgFormat, err := openImage(workingDirectory + imageID)
	if err != nil {
		fmt.Println(errors.Wrap(err, "cannot find the image file"))
	}
	imgWidth, imgHeight := src.Bounds().Max.X, src.Bounds().Max.Y
	newWidth, newHeight := 200, 200
	// new size of image
	if imgWidth > imgHeight {
		newHeight = int(float32(newHeight) * float32(imgHeight) / float32(imgWidth))
	} else {
		newWidth = int(float32(newWidth) * float32(imgWidth) / float32(imgHeight))
	}
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	// resize using given scaler
	draw.BiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	switch imgFormat {
	case "jpeg":
		jpeg.Encode(&buff, dst, nil)
	case "png":
		png.Encode(&buff, dst)
	case "gif":
		gif.Encode(&buff, dst, nil)
	}
	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	// You can embed it in an html doc with this string
	rend.HTML(w, http.StatusOK, "show", encodedString)
}

func openImage(imgFile string) (image.Image, string, error) {
	fl, err := os.Open(imgFile)
	if err != nil {
		return nil, "", errors.Wrap(err, "cannot open file")
	}
	defer fl.Close()
	img, imgType, err := image.Decode(fl)
	if err != nil {
		return nil, "", errors.Wrap(err, "cannot decode image file")
	}

	return img, imgType, nil
}
