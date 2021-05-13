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
	/*
		what does it mean the regex
		```regex
		\w=QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm
		\W=1234567890'¡`+´Ç,.-<ºª!"·$%&/()==?¿^*¨Ç;:_ŒÆÆ€®†¥  Ø∏[~§¶™ƒ∆∫Å∑©√ß µ„…–{}[[][∏Ø" ....
		```
	*/r.HandleFunc(`/getFile/{id:[\d]+}/`, controllers.GetFile)

	return http.ListenAndServe(":8080", nil)

}
