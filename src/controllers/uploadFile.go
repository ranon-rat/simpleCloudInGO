package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"github.com/ranon-rat/simpleCloudInGO/src/stuff"
	"golang.org/x/sync/errgroup"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		f, err := template.ParseFiles("public/upload.html")
		if err != nil {
			fmt.Println(err)
		}
		if err := f.Execute(w, r); err != nil {
			fmt.Println(err)
		}
		return
	default:
		if err := checkUpload(w, r); err != nil {
			log.Println(err.Err)
			http.Error(w, err.Err, err.Code)
			return
		}
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
