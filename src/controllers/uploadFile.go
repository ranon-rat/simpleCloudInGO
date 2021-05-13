package controllers

import (
	"log"
	"net/http"

	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"github.com/ranon-rat/simpleCloudInGO/src/stuff"
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
		http.ServeFile(w, r, "public/upload.html")
	case "GET":
		http.ServeFile(w, r, "public/upload.html")

	}
}

func checkUpload(w http.ResponseWriter, r *http.Request) *stuff.HttpCodeError {
	g := new(errgroup.Group)
	file, handler, err := r.FormFile("myFile")
	if err != nil {

		return nil
	}

	if handler.Size == 0 {
		log.Println("lmao this is wrong")

		return &stuff.HttpCodeError{Err: "your file is empty ", Code: 400}

	}

	g.Go(func() error { return data.UploadFile(file, handler) })
	if err := g.Wait(); err != nil {
		log.Println(err)
		return &stuff.HttpCodeError{Err: "internal server error or that file already exist", Code: 500}
	}
	return nil
}
