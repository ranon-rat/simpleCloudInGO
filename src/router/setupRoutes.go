package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/controllers"
)

func SetupRouter() error {
	r := mux.NewRouter()
	r.HandleFunc("/", nil)
	r.HandleFunc("/upload", controllers.UploadFile)
	r.HandleFunc(`/getFile/{name:[\w\W]+}`, nil)
	//	http.ServeContent(w, r, "hello world", time.Now(), bytes.NewReader([]byte("test")))
	return http.ListenAndServe(":8080", nil)

}
