package main

import (
	"os"
	"log"
	"strings"
	"image/png"
	"image/gif"
	"image/jpeg"

	"github.com/nfnt/resize"
)

/*var (
	img *image.YCbCr
)
*/
func Resize(filepath string) {
	if strings.HasSuffix(filepath, ".jpeg") || strings.HasSuffix(filepath, ".jpg") {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}

<<<<<<< HEAD
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(300, 0, img, resize.Lanczos3)
	newname := filepath + "resized.jpg"
	out, err := os.Create(newname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
=======
		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(1000, 0, img, resize.Lanczos3)
		newname := filepath[:len(filepath) - 5] + "_resized.jpg"
		out, err := os.Create(newname)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
>>>>>>> 2df2a62b0bec8c037c0ebe25bcbd5d816066d700

		// write new image to file
		jpeg.Encode(out, m, nil)

	} else if strings.HasSuffix(filepath, ".png") {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}

		// decode png into image.Image
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(1000, 0, img, resize.Lanczos3)
		newname := filepath[:len(filepath) - 4] + "_resized.png"
		out, err := os.Create(newname)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		png.Encode(out, m)

	} else if strings.HasSuffix(filepath, ".gif") {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal(err)
		}

		// decode gif into image.Image
		img, err := gif.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(1000, 0, img, resize.Lanczos3)
		newname := filepath[:len(filepath) - 4] + "_resized.gif"
		out, err := os.Create(newname)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		gif.Encode(out, m, nil)
	}
}
