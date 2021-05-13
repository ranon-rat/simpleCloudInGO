package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"github.com/ranon-rat/simpleCloudInGO/src/stuff"
)

func Api(w http.ResponseWriter, r *http.Request) {
	filesChan := make(chan []stuff.File)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	data.GetFilesName(id, filesChan)
	json.NewEncoder(w).Encode(stuff.Api{
		Files: <-filesChan,
		Size:  data.GetSize(),
	})
}
