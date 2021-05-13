package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/controllers"
)

func SetupRouter() error {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/1", http.StatusNotModified)
	})
	r.HandleFunc(`/{id:[\d]+}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "")
	})
	r.HandleFunc(`/api/{id:[\d]+}`, controllers.Api)
	r.HandleFunc("/upload", controllers.UploadFile)
	r.HandleFunc(`/getFile/{id:[\d]+}/`, controllers.GetFile)

	return http.ListenAndServe(":8080", nil)

}
