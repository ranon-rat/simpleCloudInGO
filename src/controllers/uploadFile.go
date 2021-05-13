package controllers

import (
	"log"
	"net/http"

	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"github.com/ranon-rat/simpleCloudInGO/src/stuff"
	"golang.org/x/sync/errgroup"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	case "POST":
		if err := checkUpload(w, r); err != nil {
			http.Error(w, err.Err, err.Code)
			return
		}

	default:
		w.Write([]byte("ñao ñao voce e gay"))
		return

	}
}

func checkUpload(w http.ResponseWriter, r *http.Request) *stuff.HttpCodeError {
	g := new(errgroup.Group)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println(err)
		return &stuff.HttpCodeError{Err: err.Error(), Code: 500}
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
