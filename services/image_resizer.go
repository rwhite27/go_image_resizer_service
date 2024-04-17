package services

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
)

type File struct {
	Filename  string
	Size      int64
	Extension string
}

func SaveOriginalFile(w http.ResponseWriter, r *http.Request) (string, error) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	outFile, err := os.Create(fileHeader.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return fileHeader.Filename, nil
}

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
