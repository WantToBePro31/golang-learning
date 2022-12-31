package api

import (
	"io"
	"net/http"
	"os"
	"path"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	// TODO: answer here
	http.ServeFile(w, r, "./assets/images/img-avatar.png")
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	// TODO: answer here
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	avatarFile, _, err := r.FormFile("file-avatar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer avatarFile.Close()

	fileLocation := path.Join("assets/images/img-avatar.png")
	newAvatarFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer newAvatarFile.Close()

	_, err = io.Copy(newAvatarFile, avatarFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api.dashboardView(w, r)
}
