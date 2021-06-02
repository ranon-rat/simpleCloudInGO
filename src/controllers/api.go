package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/simpleCloudInGO/src/data"
	"github.com/ranon-rat/simpleCloudInGO/src/interfaces"
)

func Api(w http.ResponseWriter, r *http.Request) {
	filesChan := make(chan []interfaces.File)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	size := data.GetSize()

	go data.GetFilesName(id, size, filesChan)

	json.NewEncoder(w).Encode(interfaces.Api{
		Files: <-filesChan,
		Size:  size,
	})
}
