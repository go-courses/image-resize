/*
 * Image Resize
 *
 * This is a Image Resize server.
 *
 * API version: 1.0.0
 * Contact: fedorenko.tolik@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	filetype "gopkg.in/h2non/filetype.v1"
)

// ImagesPost to perform resize with uploaded image
func ImagesPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer http.Redirect(w, r, "/images/100/"+handler.Filename, 302)
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
	fmt.Println("Image status is ", imageStatus, "Ваш файл прошёл проверку и успешно загружен!")

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
	//fmt.Printf("File type: %s. MIME: %s. IsImage: %v\n", kind.Extension, kind.MIME.Value, filetype.IsImage(file))
	return true, nil
}