package controllers

import (
	"log"
	"net/http"

	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"github.com/ranon-rat/simpleCloudInGO/src/interfaces"
	"golang.org/x/sync/errgroup"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	switch r.Method {
	case "POST":
		if err := checkUpload(w, r); err != nil {
			log.Println(err.Err)
			http.Error(w, err.Err, err.Code)

		}
		http.Redirect(w, r, "/1", http.StatusPermanentRedirect)
	case "GET":
		http.ServeFile(w, r, "./views/upload.html")

	}
}

func checkUpload(w http.ResponseWriter, r *http.Request) *interfaces.HttpCodeError {
	g := new(errgroup.Group)
	file, handler, err := r.FormFile("myFile")
	if err != nil {

		return nil
	}

	if handler.Size == 0 {
		log.Println("lmao this is wrong")

		return &interfaces.HttpCodeError{Err: "your file is empty ", Code: 400}

	}

	g.Go(func() error { return data.UploadFile(file, handler) })
	if err := g.Wait(); err != nil {
		log.Println(err)
		return &interfaces.HttpCodeError{Err: "internal server error or that file already exist", Code: 500}
	}
	return nil
}
