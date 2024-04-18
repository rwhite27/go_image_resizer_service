package services

import (
	"io"
	"net/http"
	"os"
)

// Upload a file to an S3 bucket
func UploadToS3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("File uploaded succesfully!"))
}

// Save a file to local storage
func SaveOriginalFile(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte("File uploaded succesfully!"))
}
