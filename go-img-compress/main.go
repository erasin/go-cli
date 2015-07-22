package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/nfnt/resize"
)

func main() {
	// open "test.jpg"
	w, _ := strconv.ParseUint(os.Args[1], 10, 32)

	// log.Println(os.Args)

	file, err := os.Open(os.Args[2])
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// info, _ := file.Stat()
	// log.Println(info.Name())

	var img image.Image

	if t, _ := regexp.MatchString(`png$`, file.Name()); t {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	if t, _ := regexp.MatchString(`jp([e]{0,1})g$`, file.Name()); t {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(w, 0, img, resize.Lanczos3)

	out, err := os.Create(os.Args[1] + "_cp.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
