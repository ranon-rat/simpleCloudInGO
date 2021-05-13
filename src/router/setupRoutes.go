package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/controllers"
)

func SetupRouter() error {
	r := mux.NewRouter().StrictSlash(true)
	/*	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/1", http.StatusNotModified)
	})*/
	r.HandleFunc(`/{id:[\d]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/index.html")
	})
	r.HandleFunc(`/api/{id:[\d]+}`, controllers.Api)
	r.HandleFunc("/upload/", controllers.UploadFile)

	r.HandleFunc(`/getFile/{id:[\d]+}/{name:[\w\W]+}/`, controllers.GetFile)
	// action="/uploadfile
	log.Println("running on localhost:8080")
	return http.ListenAndServe(":8080", r)

}
