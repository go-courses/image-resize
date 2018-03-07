package main

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

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
	}


	
}
