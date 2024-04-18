package services

import (
	"log"
	"net/http"

	"github.com/disintegration/imaging"
)

// Resize an image given the width and length
func ResizeImage(w http.ResponseWriter, filename string, width int, length int) {

	src, err := imaging.Open(filename)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	resizedImage := imaging.Resize(src, width, length, imaging.Lanczos)

	err = imaging.Save(resizedImage, "random.jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Image has been resized!"))
}
