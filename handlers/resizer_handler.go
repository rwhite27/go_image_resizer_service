package handlers

import (
	"fmt"
	"image_resizer/services"
	"net/http"
	"strconv"
)

func ResizerHandler(w http.ResponseWriter, r *http.Request) {
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

	services.ResizeImage(w, "rw_01.jpg", widthInt, lengthInt)

}
