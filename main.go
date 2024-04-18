package main

import (
	"fmt"
	"net/http"

	"image_resizer/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		filesystem := handlers.FilesystemHandler{}
		filesystem.UploadLocal(w, r)
	})

	r.Post("/resize", handlers.ResizerHandler)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
