package main

import (
	"fmt"
	"net/http"
	"strconv"

	"image_resizer/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/resize", func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()

		width := params.Get("width")
		length := params.Get("length")

		widthInt, err := strconv.Atoi(width)
		if err != nil {
			fmt.Fprint(w, "Invalid 'width' parameter.")
			return
		}

		lengthInt, err := strconv.Atoi(length)
		if err != nil {
			fmt.Fprint(w, "Invalid 'length' parameter.")
			return
		}

		// filename, err := services.SaveOriginalFile(w, r)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		services.ResizeImage(w, "rw_01.jpg", widthInt, lengthInt)

	})
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}
