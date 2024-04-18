package handlers

import (
	"image_resizer/services"
	"net/http"
)

type FilesystemHandler struct{}

func (f FilesystemHandler) UploadLocal(w http.ResponseWriter, r *http.Request) {
	services.SaveOriginalFile(w, r)
}
